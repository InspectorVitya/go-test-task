package main

import (
	"context"
	"errors"
	"github.com/inspectorvitya/go-test-task/internal/app"
	"github.com/inspectorvitya/go-test-task/internal/cache"
	"github.com/inspectorvitya/go-test-task/internal/config"
	httpserver "github.com/inspectorvitya/go-test-task/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	c, err := cache.NewCache(cfg.CapacityCache)
	if err != nil {
		log.Fatalln(err)
	}

	proxy := app.NewAppProxy(c, cfg.UrlBackend)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	server := httpserver.New(cfg, proxy)
	go func() {
		log.Println("http backend start...")
		if err := server.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("http backend http stopped....")
			} else {
				log.Fatalln(err)
			}
		}
	}()
	<-stop
	ctxClose, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = server.Stop(ctxClose)
	if err != nil {
		log.Fatalln(err)
	}
}
