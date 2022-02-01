package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-run-eas-build/step"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.NewLogger()
	step := createStep(logger)
	config, err := step.ProcessConfig()
	if err != nil {
		logger.Errorf(err.Error())
		return 1

	}

	if err := step.InstallDeps(); err != nil {
		logger.Errorf(err.Error())
		return 1
	}

	res, runErr := step.Run(config)
	exportErr := step.Export(res)

	if runErr != nil {
		logger.Errorf(runErr.Error())
		return 1
	}

	if exportErr != nil {
		logger.Errorf(exportErr.Error())
		return 1
	}

	return 0
}

func createStep(logger log.Logger) step.RunEASBuild {
	envRepository := env.NewRepository()
	inputParser := stepconf.NewInputParser(envRepository)

	return step.NewRunEASBuild(inputParser)
}
