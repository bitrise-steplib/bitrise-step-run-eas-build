package step

import (
	"fmt"
	"time"

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

	start := time.Now()

	if err := client.Build(input.Platform); err != nil {
		return fmt.Errorf("running eas build failed: %w", err)
	}

	s.logger.TDonef("Finished in: %s", runtime(time.Since(start)))

	return nil
}

func runtime(d time.Duration) string {
	const minToSec = 60
	const hToMin = 60

	elapsed := int(d / time.Second)

	min := elapsed / minToSec
	sec := elapsed - (min * minToSec)

	h := min / hToMin
	min -= h * hToMin

	s := ""
	if h > 0 {
		s += fmt.Sprintf("%dh", h)
	}
	if min > 0 {
		if len(s) > 0 {
			s += " "
		}
		s += fmt.Sprintf("%dm", min)
	}
	if sec > 0 {
		if len(s) > 0 {
			s += " "
		}
		s += fmt.Sprintf("%ds", sec)
	}
	return s
}
