package global

import (
	"github.com/johnnyeven/libtools/clients/client_id"
	"github.com/johnnyeven/libtools/courier/client"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/libtools/sqlx/mysql"
	"github.com/johnnyeven/service-configurations/database"
)

func init() {
	servicex.SetServiceName("service-configurations")
	servicex.ConfP(&Config)

	database.DBConfiguration.MustMigrateTo(Config.MasterDB.Get(), !servicex.AutoMigrate)

}

var Config = struct {
	Log    *log.Log
	Server transport_http.ServeHTTP

	MasterDB *mysql.MySQL
	SlaveDB  *mysql.MySQL

	ClientID *client_id.ClientID
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	Server: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8002,
	},

	MasterDB: &mysql.MySQL{
		Name:     "configuration",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Host:     "localhost",
	},
	SlaveDB: &mysql.MySQL{
		Name:     "configuration-readonly",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Host:     "localhost",
	},

	ClientID: &client_id.ClientID{
		Client: client.Client{
			Host: "localhost",
			Port: 8001,
		},
	},
}
