package main

import (
	"my-gomicro-client/handler"
	"net/http"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.my-gomicro-client"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/mygomicroclient/call", handler.MyGomicroClientCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
