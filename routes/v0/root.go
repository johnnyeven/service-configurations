package v0

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/service-configurations/routes/v0/configurations"
)

var Router = courier.NewRouter(V0Group{})

func init() {
	Router.Register(configurations.Router)
}

type V0Group struct {
	courier.EmptyOperator
}

func (V0Group) Path() string {
	return "/v0"
}
