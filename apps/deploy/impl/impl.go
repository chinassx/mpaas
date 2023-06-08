package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/deploy"
	cluster "github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	deploy.UnimplementedRPCServer
	ioc.IocObjectImpl

	mcenter *rpc.ClientSet
	cluster cluster.Service
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())
	i.mcenter = rpc.C()
	i.cluster = ioc.GetController(cluster.AppName).(cluster.Service)
	return nil
}

func (i *impl) Name() string {
	return deploy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	deploy.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
