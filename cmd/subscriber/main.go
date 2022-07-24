package main

import (
	"context"
	"github.com/infamax/nats-streaming-server/config"
	cache2 "github.com/infamax/nats-streaming-server/internal/cache"
	"github.com/infamax/nats-streaming-server/internal/handers"
	"github.com/infamax/nats-streaming-server/internal/repository"
	service2 "github.com/infamax/nats-streaming-server/internal/service"
	stream2 "github.com/infamax/nats-streaming-server/internal/stream"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	time.Sleep(10 * time.Second)
	f, err := os.ReadFile("config/config.yaml")

	if err != nil {
		log.Fatalf("cannot open the file: %v", err)
	}

	cf, err := config.ParseConfig(f)

	if err != nil {
		log.Fatalf("cannot parse config: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	db, err := repository.NewDB(ctx, cf.GetConnectionString())

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	cache := cache2.New()
	service, err := service2.New(db, cache)
	if err != nil {
		log.Fatalf("cannot start service: %v", err)
	}

	handler, err := handers.New(service)

	if err != nil {
		log.Fatalf("cannot start handlers: %v", err)
	}

	router := handler.InitRoutes()
	sc, err := stan.Connect(cf.SubscriberConfig.ClusterID, cf.SubscriberConfig.ClientID)
	defer sc.Close()

	if err != nil {
		log.Fatalf("cannot connect to nats streaming\n")
	}

	stream, err := stream2.New(sc, service)

	if err != nil {
		log.Fatal("cannot create nats-streaming")
	}

	stream.Start(cf.SubscriberConfig.Channel, cf.SubscriberConfig.QueueGroup)
	go func() {
		if err := router.Run("localhost:8080"); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("server start")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("server stop")
}
