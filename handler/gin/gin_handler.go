package gin

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (h *GinHandler) Start() {
	e := h.Engine.Run(":" + strconv.Itoa(h.Service.GetConfigParams().HttpPort))
	if e != nil {
		log.Fatalln(e)
	}
}

type GinHandler struct {
	Service interfaces.UseCase
	Engine  *gin.Engine
}
