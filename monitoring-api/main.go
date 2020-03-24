package main

import (
	"github.com/Dadard29/podman-monitoring/monitoring-api/api"
	"os"
)

func main() {
	a := api.NewApi(
		api.ServerConfig{
			Host: "0.0.0.0",
			Port: "9000",
		},
		api.DbConfig{
			Username:     os.Getenv("USERNAME_DB"),
			Password:     os.Getenv("PASSWORD_DB"),
			Host:         "127.0.0.1",
			Port:         "3306",
			DatabaseName: "monitoring",
		})

	a.Serve()
	a.Stop()
}
