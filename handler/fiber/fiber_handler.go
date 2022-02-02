package fiber

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (h *FiberHandler) Start() {
	app := fiber.New(fiber.Config{})

	InitRoutes(app, h)

	e := h.Engine.Run(":" + strconv.Itoa(h.HUseCase.GetConfigParams().HttpPort))
	if e != nil {
		log.Fatalln(e)
	}
}

type FiberHandler struct {
	HUseCase interfaces.UseCase
	Engine   *gin.Engine
}

func (h *FiberHandler) UseCase() interfaces.UseCase {
	return h.HUseCase
}

func InitRoutes(app *fiber.App, h interfaces.Handler) {
	group := app.Group("/")

	errorRoutes(group, h)
}

func errorRoutes(r fiber.Router, h interfaces.Handler) {
	er := &ErrorRouter{
		Handler: h,
	}

	r.Get("/error/:id", er.GetErrorById)
	r.Post("/error", er.AddError)
	r.Get("/errors", er.GetAllErrors)
	r.Delete("error/:id", er.DeleteError)
	r.Put("/error/:id", er.EditError)
}
