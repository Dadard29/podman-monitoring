package api

import (
	"encoding/json"
	"github.com/Dadard29/podman-monitoring/scraper"
	"io/ioutil"
	"log"
	"net/http"
)

func storePodInfos(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var jsonBody []scraper.PodInfos
	err = json.Unmarshal(data, &jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = storePodInfosRepo(jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func storePodmanInfos(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var jsonBody scraper.PodmanInfos
	err = json.Unmarshal(data, &jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = storePodmanInfosRepo(jsonBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func storePodmanProxyInfos(w http.ResponseWriter, r *http.Request) {
	// todo
}
