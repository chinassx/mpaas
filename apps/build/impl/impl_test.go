package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mpaas/apps/build"
	"github.com/infraboard/mpaas/test/tools"
)

var (
	impl build.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(build.AppName).(build.Service)
}
