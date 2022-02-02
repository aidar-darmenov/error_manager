package interfaces

import (
	"github.com/aidar-darmenov/error_manager/config"
)

type Configuration interface {
	InitConfigParams(string)
	Params() *config.Configuration
}
