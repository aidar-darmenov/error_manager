package fiber

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func (h *FiberHandler) Start() {
	app := fiber.New(fiber.Config{})

	InitRoutes(app, h)

}

type FiberHandler struct {
	HUseCase   interfaces.UseCase
	Router     fiber.Router
	GrpcRouter *grpc.Server
}

func (h *FiberHandler) UseCase() interfaces.UseCase {
	return h.HUseCase
}

func InitRoutes(app *fiber.App, h interfaces.Handler) {
	group := app.Group("/")

	errorRoutes(group, h)
}

func errorRoutes(r fiber.Router, h interfaces.Handler) {

	r.Get("/error/:id", h.GetErrorById)
	r.Post("/error", h.AddError)
	r.Get("/errors", h.GetAllErrors)
	r.Delete("error/:id", h.DeleteError)
	r.Put("/error/:id", h.EditError)
}
