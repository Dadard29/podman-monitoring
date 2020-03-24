package main

import (
	"github.com/Dadard29/podman-monitoring/api"
	"github.com/Dadard29/podman-monitoring/scraper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func task(t time.Time) {
	log.Println(t)
}

func scrape() {
	apiIp := os.Getenv("API_HOST")
	s := scraper.NewScraper(apiIp)

	// use ticker to repeat the task every n seconds
	n := 5
	log.Println("setting up ticker with period", n)
	tick := time.NewTicker(time.Second * time.Duration(n))
	done := make(chan bool)
	go func(tick *time.Ticker, done chan bool) {
		s.MainTask(time.Now())
		for {
			select {
			case t := <-tick.C:
				s.MainTask(t)
			case <-done:
				return
			}
		}
	}(tick, done)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	done <- true
}

func main() {
	cmd := os.Args[1]
	if cmd == "scraper" {
		scrape()
	} else if cmd == "api" {
		serve()
	}
}
