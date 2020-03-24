package api

import (
	"encoding/json"
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

	var jsonBody []PodInfosJson
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

	var jsonBody PodmanInfosJson
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
