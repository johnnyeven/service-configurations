package configurations

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-configurations/global"
	"github.com/johnnyeven/service-configurations/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(UpdateConfiguration{}))
}

// 更新配置
type UpdateConfiguration struct {
	httpx.MethodPatch
	ConfigurationID uint64                          `name:"configID,string" in:"path" default:""`
	Body            modules.UpdateConfigurationBody `name:"body" in:"body"`
}

func (req UpdateConfiguration) Path() string {
	return "/:configID"
}

func (req UpdateConfiguration) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.UpdateConfiguration(req.ConfigurationID, req.Body, db)
	if err != nil {
		logrus.Errorf("[UpdateConfiguration] modules.UpdateConfiguration err: %v, request: %+v", err, req)
	}
	return
}
