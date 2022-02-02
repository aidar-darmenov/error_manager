package interfaces

import (
	"github.com/aidar-darmenov/error_manager/config"
	"go.uber.org/zap"
)

type UseCase interface {
	GetLogger() *zap.Logger
	GetConfigParams() *config.Configuration
}
