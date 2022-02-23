package step

import (
	"fmt"
	"time"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/kballard/go-shellquote"
)

type EASClientBuilder interface {
	Build(token stepconf.Secret, workDir string) EASClient
}

type EASClient interface {
	Build(platform string, options ...string) error
}

type Input struct {
	AccessToken stepconf.Secret `env:"access_token,required"`
	Platform    string          `env:"platform,opt[all,android,ios]"`
	WorkDir     string          `env:"work_dir,dir"`
	EASOptions  string          `env:"eas_options"`
}

type Config struct {
	AccessToken stepconf.Secret
	Platform    string
	WorkDir     string
	EASOptions  []string
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

func (s EASBuilder) ProcessConfig() (Config, error) {
	var input Input
	err := s.inputParser.Parse(&input)
	if err != nil {
		return Config{}, err
	}
	stepconf.Print(input)

	var options []string
	if len(input.EASOptions) > 0 {
		var err error
		options, err = shellquote.Split(input.EASOptions)
		if err != nil {
			return Config{}, fmt.Errorf("eas_options are not valid CLI arguments: %w", err)
		}
	}

	return Config{
		AccessToken: input.AccessToken,
		Platform:    input.Platform,
		WorkDir:     input.WorkDir,
		EASOptions:  options,
	}, nil
}

func (s EASBuilder) Run(cfg Config) error {
	client := s.clientBuilder.Build(cfg.AccessToken, cfg.WorkDir)

	s.logger.Println()
	s.logger.TInfof("Running EAS build")

	start := time.Now()

	if err := client.Build(cfg.Platform, cfg.EASOptions...); err != nil {
		return fmt.Errorf("running eas build failed: %w", err)
	}

	s.logger.TDonef("Finished in: %s", runtime(time.Since(start)))

	return nil
}

func runtime(d time.Duration) string {
	return fmt.Sprint(d.Round(time.Second))
}
