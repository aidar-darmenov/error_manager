package fiber

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gofiber/fiber/v2"
)

type ErrorRouter struct {
	Handler interfaces.Handler
}

func (r *FiberHandler) GetErrorById(c *fiber.Ctx) error {
	return nil
}
func (r *FiberHandler) AddError(c *fiber.Ctx) error {
	return nil
}
func (r *FiberHandler) GetAllErrors(c *fiber.Ctx) error {
	return nil
}
func (r *FiberHandler) DeleteError(c *fiber.Ctx) error {
	return nil
}
func (r *FiberHandler) EditError(c *fiber.Ctx) error {
	return nil
}
