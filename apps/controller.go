package apps

import (
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/infraboard/mpaas/apps/approval/impl"
	_ "github.com/infraboard/mpaas/apps/build/impl"
	_ "github.com/infraboard/mpaas/apps/cluster/impl"
	_ "github.com/infraboard/mpaas/apps/deploy/impl"
	_ "github.com/infraboard/mpaas/apps/gateway/impl"
	_ "github.com/infraboard/mpaas/apps/job/impl"
	_ "github.com/infraboard/mpaas/apps/k8s/impl"
	_ "github.com/infraboard/mpaas/apps/log/impl"
	_ "github.com/infraboard/mpaas/apps/pipeline/impl"
	_ "github.com/infraboard/mpaas/apps/task/impl"
	_ "github.com/infraboard/mpaas/apps/trigger/impl"
)
