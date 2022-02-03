package interfaces

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Start()
	UseCase() UseCase
	ErrorHandler
}

type ErrorHandler interface {
	GetErrorById(*fiber.Ctx) error
	AddError(*fiber.Ctx) error
	GetAllErrors(*fiber.Ctx) error
	DeleteError(*fiber.Ctx) error
	EditError(*fiber.Ctx) error
}

type GrpcRouter interface {
}
