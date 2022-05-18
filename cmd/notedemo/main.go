package main

import (
	"easy_note/cmd/notedemo/dal"
	notedemo "easy_note/cmd/notedemo/kitex_gen/notedemo/noteservice"
	"easy_note/cmd/notedemo/rpc"
	"easy_note/pkg/bound"
	"easy_note/pkg/constants"
	"easy_note/pkg/middleware"
	"easy_note/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

func Init() {
	tracer.InitJaeger(constants.NoteServiceName)
	rpc.InitRPC()
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	Init()

	svr := notedemo.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.NoteServiceName}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
		server.WithBoundHandler(bound.NewCpuLimitHandler()),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
