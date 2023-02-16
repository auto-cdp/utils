module github.com/glory-cd/utils

go 1.12

require (
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/containerd/containerd v1.5.18 // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/coreos/go-systemd v0.0.0-20190620071333-e64a0ec8b42a // indirect
	github.com/docker/docker v0.0.0-20190813234819-fade624f1696
	github.com/docker/go-connections v0.4.0
	github.com/garyburd/redigo v1.6.0
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/hashicorp/go-getter v1.3.0
	github.com/jinzhu/gorm v1.9.10 // indirect
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c // indirect
	github.com/ory/dockertest v3.3.4+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/wantedly/gorm-zap v0.0.0-20171015071652-372d3517a876
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.5.0
	google.golang.org/grpc/examples v0.0.0-20230215194445-0f02ca5cc927 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/ory-am/dockertest.v3 v3.3.4 // indirect
)

replace github.com/glory-cd/utils => ./
