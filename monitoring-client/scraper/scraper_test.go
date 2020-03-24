package scraper

import "testing"

func TestScraper_GetAndSendPodmanInfos(t *testing.T) {
	s := NewScraper("ip")
	s.GetAndSendPodmanInfos()
}
