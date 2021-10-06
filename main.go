package main

import (
	"context"

	_ "github.com/lib/pq"

	"os"
	"os/signal"
	"restapi/pkg/router"

	"github.com/rs/zerolog/log"
)

func main() {

	//Gracefull shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()
	if err := router.Router(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}

}
