package main

import (
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-configurations/global"
	"github.com/johnnyeven/service-configurations/routes"
)

func main() {
	servicex.Execute()
	global.Config.Server.Serve(routes.RootRouter)
}
