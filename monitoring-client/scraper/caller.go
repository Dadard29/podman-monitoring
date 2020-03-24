package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	podRoute         = "/pods"
	podmanInfosRoute = "/podman/infos"
	podmanProxyRoute = "/podman/proxy"
)

func (s Scraper) sendPodInfos(podInfos []PodInfos) error {
	data, err := json.Marshal(&podInfos)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", s.apiHost, podRoute)
	buf := bytes.NewBuffer(data)
	r, err := http.Post(url, "application/json", buf)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status is %s", r.Status))
	}

	return nil
}

func (s Scraper) sendPodmanInfos(infos PodmanInfos) error {
	data, err := json.Marshal(&infos)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s", s.apiHost, podmanInfosRoute)
	buf := bytes.NewBuffer(data)
	r, err := http.Post(url, "application/json", buf)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status is %s", r.Status))
	}

	return nil
}
