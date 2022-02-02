package fiber

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/gofiber/fiber/v2"
)

type ErrorRouter struct {
	Handler interfaces.Handler
}

func (r *ErrorRouter) GetErrorById(c *fiber.Ctx) error {
	return nil
}
func (r *ErrorRouter) AddError(c *fiber.Ctx) error {
	return nil
}
func (r *ErrorRouter) GetAllErrors(c *fiber.Ctx) error {
	return nil
}
func (r *ErrorRouter) DeleteError(c *fiber.Ctx) error {
	return nil
}
func (r *ErrorRouter) EditError(c *fiber.Ctx) error {
	return nil
}
