package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/linger1216/go-front/echo-service/svc"
)

func CustomizedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// customized middleware implement
	}
}

func WrapEndpoints(in svc.Endpoints) svc.Endpoints {
	return in
}

func WrapService(in svc.EchoServer) svc.EchoServer {
	return in
}
