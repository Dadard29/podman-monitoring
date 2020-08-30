package api

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func (a Api) Serve() {
	go func() {
		log.Println("serving on ", a.server.Addr)
		a.server.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (a Api) Stop() {
	log.Println("closing service")
	var wait time.Duration
	ctx, _ := context.WithTimeout(context.Background(), wait)
	a.server.Shutdown(ctx)
}
