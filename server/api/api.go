package api

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

var apiObject Api

type Api struct {
	server *http.Server
	orm    *gorm.DB
}

type DbConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

type ServerConfig struct {
	Host string
	Port string
}

func newServer(serverConfig ServerConfig) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/pods", storePodInfos).Methods(http.MethodPost)
	r.HandleFunc("/podman/infos", storePodmanInfos).Methods(http.MethodPost)
	r.HandleFunc("/podman/proxy", storePodmanProxyInfos).Methods(http.MethodPost)

	return &http.Server{
		Addr:              fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port),
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          log.New(
			os.Stdout, "podman-monitoring-server", log.Ldate|log.Ltime),
		BaseContext:       nil,
		ConnContext:       nil,
	}
}

func newOrm(config DbConfig) *gorm.DB {
	parseTime := "?parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		config.Username, config.Password,
		config.Host, config.Port,
		config.DatabaseName, parseTime)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func NewApi(serverConfig ServerConfig, dbConfig DbConfig) Api {
	server := newServer(serverConfig)
	orm := newOrm(dbConfig)

	apiObject = Api{
		server: server,
		orm:    orm,
	}

	return apiObject
}
