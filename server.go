package main

import (
	"github.com/aidar-darmenov/error_manager/config"
	"github.com/aidar-darmenov/error_manager/handler"
	"github.com/aidar-darmenov/error_manager/repository"
	"github.com/aidar-darmenov/error_manager/usecase"
	"go.uber.org/zap"
	"log"
)

func main() {

	cfg := config.NewConfiguration("config/config.json")

	// Used uber zap logger for simple example. Now it writes in console
	// Usually, for this purposes we use logs sent to Kibana Elastic Search through Kafka
	var loggerConfig = zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(zap.DebugLevel)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	r := repository.NewRepository()

	// Creating abstract usecase(business logic) layer
	u := usecase.NewUseCase(cfg, logger, r)

	// Creating abstract webService(delivery) layer
	h := handler.NewHandler(u)
	h.Start()
}
