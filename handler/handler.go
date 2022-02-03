package handler

import (
	"github.com/aidar-darmenov/error_manager/handler/fiber"
	"github.com/aidar-darmenov/error_manager/handler/grpc"
	"github.com/aidar-darmenov/error_manager/interfaces"
)

func NewHandler(u interfaces.UseCase) interfaces.Handler {

	h := &fiber.FiberHandler{
		HUseCase:   u,
		GrpcRouter: grpc.InitRouter(*u.GetConfigParams()),
	}

	return h
}
