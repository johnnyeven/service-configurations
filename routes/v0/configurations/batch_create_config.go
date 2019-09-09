package configurations

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/service-configurations/global"
	"github.com/johnnyeven/service-configurations/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(BatchCreateConfig{}))
}

// 批量创建配置
type BatchCreateConfig struct {
	httpx.MethodPost
	Body []modules.CreateConfigurationBody `name:"body" in:"body"`
}

func (req BatchCreateConfig) Path() string {
	return "/:configID/batch"
}

func (req BatchCreateConfig) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	tx := sqlx.NewTasks(db)

	for _, param := range req.Body {
		tx = tx.With(GetTransaction(param))
	}

	if err = tx.Do(); err != nil {
		logrus.Errorf("[BatchCreateConfig] transaction modules.CreateConfiguration err: %v, req: %+v", err, req.Body)
	}
	return
}

func GetTransaction(req modules.CreateConfigurationBody) func(db *sqlx.DB) error {
	return func(db *sqlx.DB) error {
		err := modules.CreateConfiguration(req, db, global.Config.ClientID)
		if err != nil {
			return err
		}
		return nil
	}
}
