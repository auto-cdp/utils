package afis

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"io"
	"os"
	"testing"
)

func TestDownloadNormal(t *testing.T) {

	cmd := strslice.StrSlice{"--auth-type", "http", "--auth-http", "admin:afis2020", "--upload"}
	ca := ContainerAttr{}
	ca.image = "liupzmin/httpfileserver:2.0"
	ca.cport = "8000"
	ca.hport = "8000"
	ca.cmd = cmd
	ca.cname = "httpserver"

	c, err := createContainer(ca)

	if err != nil{
		t.Errorf(`createContainer(create container failed: %s) = error`, err.Error())
	}
	err = DownloadCode("/tmp", "http://admin:afis2020@127.0.0.1:8000/Dockerfile")

	if err != nil {
		t.Errorf(`"/tmp", "http://admin:afis2020@127.0.0.1:8000/Dockerfile %s") = error`, err.Error())
	}

	err = stopContainer(c)

	if err != nil{
		t.Errorf(`createContainer(stop container failed: %s) = error`, err.Error())
	}

	err = removeContainer(c)

	if err != nil{
		t.Errorf(`createContainer(remove container failed: %s) = error`, err.Error())
	}
}


func TestDownloadNoNormal(t *testing.T) {
	err := DownloadCode("/root/test", "http://lp:0124@127.0.0.1/hfp/tst.tar.gz")

	if err == nil {
		t.Error(`Download("/root/test:http://lp:0124@127.0.0.1/hfp/tst.tar.gz") = true`)
	}

	if err == nil {
		t.Error(`Download("/home/liupeng/test:http://lp:012w4@127.0.0.1/hfp/tst.tar.gz") = true`)
	}
}

func TestCheckFileOwnerNormal(t *testing.T) {
	if !CheckFileOwner("/home/liupeng/AlarmWechat", "liupeng") {
		t.Error("/home/liupeng/AlarmWechat:liupeng = error")
	}
}

func TestCheckFileOwnerNoNormal(t *testing.T) {
	if CheckFileOwner("/home/liupeng/AlarmWechat", "alarm") {
		t.Error("/home/liupeng/AlarmWechat:alarm = error")
	}
}

func TestIsExecutableNormal(t *testing.T) {
	if !IsExecutable("/home/liupeng/ant") {
		t.Error("/home/liupeng/ant = error")
	}
}

func TestIsExecutableNoNormal(t *testing.T) {
	if IsExecutable("/home/liupeng/tcpdump.cap") {
		t.Error("/home/liupeng/tcpdump.cap = error")
	}
}

func TestIntegerStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

type ContainerAttr struct {
	image string
	cname string
	cport string
	hport string
	cmd   strslice.StrSlice
}

func createContainer(conattr ContainerAttr) (string, error) {

	cli, err := initDockerClient()

	if err != nil {
		return "", err
	}

	ctx := context.Background()
	imageName := conattr.image

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}

	_, err = io.Copy(os.Stdout, out)

	if err != nil {
		return "", err
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: conattr.hport,
	}
	exports := make(nat.PortSet, 10)
	containerPort, err := nat.NewPort("tcp", conattr.cport)
	exports[containerPort] = struct{}{}
	if err != nil {
		return "", err
	}

	portMap := make(nat.PortMap, 0)
	tmp := make([]nat.PortBinding, 0, 1)
	tmp = append(tmp, hostBinding)
	portMap[containerPort] = tmp

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,
		ExposedPorts: exports,
		Cmd:          conattr.cmd,
	},
		&container.HostConfig{
			PortBindings: portMap,
		}, nil, "")

	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}

func initDockerClient() (*client.Client, error) {


	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli, nil
}


func stopContainer(containerID string) error{
	cli, err := initDockerClient()

	if err != nil {
		return err
	}

	ctx := context.Background()

	if err = cli.ContainerStop(ctx, containerID, nil); err != nil {
		return err
	}

	return nil
}

func removeContainer(containerID string) error{
	cli, err := initDockerClient()

	if err != nil {
		return err
	}

	ctx := context.Background()

	if err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{}); err != nil {
		return err
	}

	return nil
}