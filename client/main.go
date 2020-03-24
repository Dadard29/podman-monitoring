package main

import (
	"github.com/Dadard29/podman-monitoring/client/scraper"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	apiIp := os.Getenv("API_HOST")
	s := scraper.NewScraper(apiIp)

	// use ticker to repeat the task every n seconds
	n, err := strconv.Atoi(os.Getenv("SCRAPE_PERIOD"))
	if err != nil {
		log.Fatalln("error parsing env SCRAPE_PERIOD", err)
		return
	}
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
