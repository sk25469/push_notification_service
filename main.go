package main

import (
	"github.com/sk25469/push_noti_service/pkg/config"
	"github.com/sk25469/push_noti_service/pkg/routes"
)

func main() {

	config.InitPostgres()
	routes.InitRoutes()

}
