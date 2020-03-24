package scraper

import (
	"encoding/json"
	"fmt"
)

const (
	podRoute           = "/pods"
	podmanVersionRoute = "/podman/infos"
	podmanProxyRoute   = "/podman/proxy"
)

func (s Scraper) sendPodInfos(podInfos []PodInfos) error {
	data, err := json.Marshal(&podInfos)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func (s Scraper) sendPodmanVersion(version PodmanInfos) error {
	return nil
}
