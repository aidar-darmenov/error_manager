package usecase

import (
	"github.com/aidar-darmenov/error_manager/config"
	"github.com/aidar-darmenov/error_manager/interfaces"
	"go.uber.org/zap"
)

type UseCase struct {
	Configuration interfaces.Configuration
	Logger        *zap.Logger
	Repository    interfaces.Repository
}

func NewUseCase(cfg *config.Configuration, logger *zap.Logger, repository interfaces.Repository) interfaces.UseCase {
	//Here can be any other objects like DB, Cache, any kind of delivery services

	return &UseCase{
		Configuration: cfg,
		Logger:        logger,
		Repository:    repository,
	}
}

func (s *UseCase) GetLogger() *zap.Logger {
	return s.Logger
}

func (s *UseCase) GetConfigParams() *config.Configuration {
	return s.Configuration.Params()
}
