package main

import (
	handlers "github.com/NKV510/pproject/pkg/Handlers"
	"github.com/NKV510/pproject/pkg/worker"
	"github.com/NKV510/pproject/server"
)

func main() {
	worker := worker.NewList()
	httpHandlers := handlers.NewHTTPHandlers(worker)
	httpServer := server.NewHTTPServer(httpHandlers)

	if err := httpServer.HTTPServerStart(); err != nil {
		panic(err.Error())
	}

}
