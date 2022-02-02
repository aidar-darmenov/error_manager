package handler

import (
	gin_handler "github.com/aidar-darmenov/error_manager/handler/gin"
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gin-gonic/gin"
)

func NewHandler(u interfaces.UseCase) interfaces.Handler {
	g := gin.Default()

	h := &gin_handler.GinHandler{
		Service: u,
		Engine:  g,
	}

	return h
}
