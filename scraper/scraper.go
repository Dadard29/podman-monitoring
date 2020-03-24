package scraper

import (
	"context"
	"github.com/containers/libpod/libpod"
	"log"
	"net/http"
	"time"
)

type Scraper struct {
	libpodRuntime *libpod.Runtime
	client        *http.Client
	apiHost         string
}

func NewScraper(apiHost string) Scraper {
	c := context.Background()
	runtime, err := libpod.NewRuntime(c)
	if err != nil {
		log.Fatalln(err)
	}

	return Scraper{
		libpodRuntime: runtime,
		client:        &http.Client{},
		apiHost:         apiHost,
	}
}

func (s Scraper) MainTask(t time.Time) {
	log.Println("sending infos...")
	s.GetAndSendPodmanInfos()
	s.GetAndSendPodInfos()
}

func (s Scraper) GetAndSendPodInfos() {
	podListInfos, err := s.listPodsInfos()
	if err != nil {
		log.Println(err)
		return
	}

	err = s.sendPodInfos(podListInfos)
	if err != nil {
		log.Println(err)
	}
}

func (s Scraper) GetAndSendPodmanInfos() {
	infos, err := s.getPodmanVersion()
	if err != nil {
		log.Println(err)
		return
	}

	err = s.sendPodmanInfos(infos)
	if err != nil {
		log.Println(err)
	}
}
