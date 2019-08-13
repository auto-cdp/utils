package afis

import "testing"

func TestDownloadNormal(t *testing.T) {
	err := DownloadCode("/home/liupeng/test", "http://lp:0124@127.0.0.1/hfp/tst.tar.gz?archive=false")

	if err != nil {
		t.Errorf(`Download("/home/liupeng/test:http://lp:0124@127.0.0.1/hfp/tst.tar.gz?archive=false %s") = error`, err.Error())
	}
}

func TestDownloadNormalFtp(t *testing.T) {
	err := DownloadCode("/home/liupeng/test", "ftp://devops:afis2017@192.168.1.9/code/test/deliver_release.py")

	if err != nil {
		t.Errorf(`Download("/home/liupeng/test:ftp://192.168.1.9/code/test/deliver_release.py %s") = error`, err.Error())
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


func TestCheckFileOwnerNormal(t *testing.T){
	if !CheckFileOwner("/home/liupeng/AlarmWechat", "liupeng"){
		t.Error("/home/liupeng/AlarmWechat:liupeng = error")
	}
}

func TestCheckFileOwnerNoNormal(t *testing.T){
	if CheckFileOwner("/home/liupeng/AlarmWechat", "alarm"){
		t.Error("/home/liupeng/AlarmWechat:alarm = error")
	}
}

func TestIsExecutableNormal(t *testing.T) {
	if !IsExecutable("/home/liupeng/ant"){
		t.Error("/home/liupeng/ant = error")
	}
}

func TestIsExecutableNoNormal(t *testing.T) {
	if IsExecutable("/home/liupeng/tcpdump.cap"){
		t.Error("/home/liupeng/tcpdump.cap = error")
	}
}