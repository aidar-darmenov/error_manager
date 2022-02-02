package gin

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (ws *GinHandler) Start() {
	e := ws.Engine.Run(":" + strconv.Itoa(ws.Service.GetConfigParams().HttpPort))
	if e != nil {
		log.Fatalln(e)
	}
}

type GinHandler struct {
	Service interfaces.UseCase
	Engine  *gin.Engine
}
