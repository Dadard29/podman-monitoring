package main

import (
	"github.com/Dadard29/podman-monitoring/api"
	"github.com/Dadard29/podman-monitoring/scraper"
	"os"
)

func serve() {
	a := api.NewApi(
		api.ServerConfig{
			Host: "0.0.0.0",
			Port: "9000",
		},
		api.DbConfig{
			Username:     os.Getenv("USERNAME_DB"),
			Password:     os.Getenv("PASSWORD_DB"),
			Host:         "localhost",
			Port:         "3306",
			DatabaseName: "monitoring",
		})

	a.Serve()
	a.Stop()
}

func scrape() {
	apiIp := os.Getenv("API_IP")
	s := scraper.NewScraper(apiIp)

	// ======= PODS ======= \\
	s.GetAndSendPodInfos()

	// ======= PODMAN VERSION ======= \\
	s.GetAndSendPodmanInfos()

	// ======= PODMAN-PROXY ======= \\
	// TODO
}

func main() {
	cmd := os.Args[1]
	if cmd == "scraper" {
		scrape()
	} else if cmd == "api" {
		serve()
	}
}
