package bound

import (
	"context"
	"easy_note/pkg/constants"
	"easy_note/pkg/errno"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/shirou/gopsutil/cpu"
	"net"
)

type cpuLimitHandler struct {
}

func NewCpuLimitHandler() remote.InboundHandler {
	return &cpuLimitHandler{}
}

func (c cpuLimitHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	return ctx, nil
}

func (c cpuLimitHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	return ctx
}

func (c cpuLimitHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	percent, _ := cpu.Percent(0, false)
	p := percent[0]
	klog.CtxInfof(ctx, "current cpu is %.2g", p)
	if p > constants.CPURateLimit {
		return ctx, errno.ServiceErr.WithMessage(fmt.Sprintf("cpu = %.2g", c))
	}
	return ctx, nil
}

func (c cpuLimitHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	return ctx, nil
}
