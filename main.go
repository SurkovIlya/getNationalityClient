package main

import (
	"context"
	"getNationalClient/internal/configs"
	"getNationalClient/internal/exception"
	"getNationalClient/internal/nationalpredict"
	"getNationalClient/internal/nationalsource"
	"getNationalClient/internal/service"
	"getNationalClient/internal/service/handler"
	"getNationalClient/internal/service/server"
	"getNationalClient/pkg/cache"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// const host = "https://api.nationalize.io"
// const port = "8080"

func main() {
	cfg := configs.New()
	ns := nationalsource.New(cfg.NationalData.Host)

	cl, err := nationalpredict.GetCountryList()
	if err != nil {
		log.Panic("CountryList error:", err)
	}

	np := nationalpredict.New(cl, ns)

	exc := exception.New()
	cache := cache.NewCash(uint32(cfg.Cache.TimeToLifeCache), cfg.Cache.Count)

	sv := service.New(np, exc, cache)

	handlers := handler.NewHandler(sv)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}

	}()

	log.Print("NationalServer Started")
	cfg.Print()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("NationalServer Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Panicf("error occured on server shutting down: %s", err.Error())
	}

}
