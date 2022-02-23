package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-run-eas-build/step"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()

	s := createStep(logger)
	cfg, err := s.ProcessConfig()
	if err != nil {
		logger.Errorf("Process config: %s", err)
		return 1

	}

	if err := s.Run(cfg); err != nil {
		logger.Errorf("Run: %s", err)
		return 1
	}

	return 0
}

func createStep(logger log.Logger) step.EASBuilder {
	envRepository := env.NewRepository()
	inputParser := stepconf.NewInputParser(envRepository)
	commandFactory := command.NewFactory(envRepository)
	clientBuilder := step.NewEASClientBuilder(commandFactory, logger)

	return step.NewEASBuilder(inputParser, logger, clientBuilder)
}
