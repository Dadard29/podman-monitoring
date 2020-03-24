package scraper

import (
	"context"
	"github.com/containers/libpod/libpod"
	"log"
	"net/http"
)

type Scraper struct {
	libpodRuntime *libpod.Runtime
	client        *http.Client
	apiIP         string
}

func NewScraper(apiIP string) Scraper {
	c := context.Background()
	runtime, err := libpod.NewRuntime(c)
	if err != nil {
		log.Fatalln(err)
	}

	return Scraper{
		libpodRuntime: runtime,
		client:        &http.Client{},
		apiIP:         apiIP,
	}
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

	err = s.sendPodmanVersion(infos)
	if err != nil {
		log.Println(err)
	}
}
