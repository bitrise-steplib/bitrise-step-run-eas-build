package step

import (
	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-run-eas-build/eas"
)

type easClientBuilder struct {
	commandFactory command.Factory
	logger         log.Logger
}

func NewEASClientBuilder(commandFactory command.Factory, logger log.Logger) EASClientBuilder {
	return easClientBuilder{
		commandFactory: commandFactory,
		logger:         logger,
	}
}

func (b easClientBuilder) Build(token stepconf.Secret, workDir string) EASClient {
	return eas.NewClient(b.commandFactory, b.logger, token, workDir)
}
