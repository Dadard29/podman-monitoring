package main

import (
	"fmt"
	"github.com/Dadard29/podman-monitoring/monitoring-client/scraper"
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
	varName := "SCRAPE_PERIOD"
	n, err := strconv.Atoi(os.Getenv(varName))
	if err != nil {
		log.Fatalln(fmt.Sprintf("error parsing env %s", varName), err)
		return
	}
	log.Println("setting up ticker with period", n)
	tick := time.NewTicker(time.Second * time.Duration(n))
	done := make(chan bool)

	// the podman infos are sent every k iterations of pods infos
	go func(tick *time.Ticker, done chan bool) {
		s.MainTask(time.Now())

		initialK := 3
		k := initialK
		for {
			select {
			case t := <-tick.C:
				log.Println("tick", t)
				if k == 0 {
					log.Println("sending podman infos")
					s.GetAndSendPodmanInfos()
					k = initialK
				}
				log.Println("sending pods infos")
				s.GetAndSendPodInfos()
				k -= 1
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
