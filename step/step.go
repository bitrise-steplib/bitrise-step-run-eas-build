package step

import (
	"fmt"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/log"
)

type EASClientBuilder interface {
	Build(token stepconf.Secret, workDir string) EASClient
}

type EASClient interface {
	Build(platform string) error
}

type Input struct {
	Token    stepconf.Secret `env:"access_token,required"`
	Platform string          `env:"platform,opt[all,android,ios]"`
	WorkDir  string          `env:"work_dir,dir"`
}

type EASBuilder struct {
	inputParser   stepconf.InputParser
	logger        log.Logger
	clientBuilder EASClientBuilder
}

func NewEASBuilder(inputParser stepconf.InputParser, logger log.Logger, clientBuilder EASClientBuilder) EASBuilder {
	return EASBuilder{
		inputParser:   inputParser,
		logger:        logger,
		clientBuilder: clientBuilder,
	}
}

func (s EASBuilder) ProcessConfig() (Input, error) {
	var input Input
	err := s.inputParser.Parse(&input)
	if err != nil {
		return Input{}, err
	}
	stepconf.Print(input)

	return input, nil
}

func (s EASBuilder) Run(input Input) error {
	client := s.clientBuilder.Build(input.Token, input.WorkDir)

	s.logger.Println()
	s.logger.TInfof("Running EAS build")

	if err := client.Build(input.Platform); err != nil {
		return fmt.Errorf("running eas build failed: %w", err)
	}

	return nil
}
