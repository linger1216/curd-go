package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/linger1216/go-front/service/handlers"
	"github.com/linger1216/go-front/service/svc"
	"github.com/linger1216/go-front/utils/config"
	"os"
	"os/signal"
	"syscall"
)

func NewEndpoints(reader config.Reader) svc.Endpoints {
	service := handlers.NewService(reader)
	service = handlers.WrapService(service)
	endpoints := svc.Endpoints{
		CreateEchoEndpoint: svc.MakeCreateEchoEndpoint(service),
		GetEchoEndpoint:    svc.MakeGetEchoEndpoint(service),
		ListEchoEndpoint:   svc.MakeListEchoEndpoint(service),
		UpdateEchoEndpoint: svc.MakeUpdateEchoEndpoint(service),
		DeleteEchoEndpoint: svc.MakeDeleteEchoEndpoint(service),
	}
	endpoints = handlers.WrapEndpoints(endpoints)
	return endpoints
}

func interruptHandler(ch chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)
	ch <- terminateError
}

func Run(reader config.Reader) {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	engine.Use(handlers.CustomizedMiddleware())

	ch := make(chan error)
	go interruptHandler(ch)

	endpoints := NewEndpoints(reader)

	// Debug listener.
	go func() {
		addr := reader.GetString("server", "debugAddr")
		e := svc.MakeDebugHandler(addr)
		ch <- e.Run(addr)
	}()

	// http
	go func() {
		addr := reader.GetString("server", "httpAddr")
		svc.MakeHTTPHandler(engine, endpoints)
		ch <- engine.Run(addr)
	}()

	fmt.Printf("echo server started.\n")
	fmt.Printf("closed:%s", <-ch)
}
