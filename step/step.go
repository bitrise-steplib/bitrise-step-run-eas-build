package step

import (
	"errors"
	"strings"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
)

// Input ...
type Input struct {
	User      stepconf.Secret `env:"user,required"`
	Password  stepconf.Secret `env:"password,required"`
	Platforms string          `env:"platforms,required"`
}

// Config ...
type Config struct {
	User      stepconf.Secret
	Password  stepconf.Secret
	Platforms []string
}

type ExpoClient interface {
	Login(user, password stepconf.Secret) error
	Build(platforms []string, options ...string) error
}

// RunEASBuild ...
type RunEASBuild struct {
	inputParser stepconf.InputParser
	expoClient  ExpoClient
}

// NewRunEASBuild ...
func NewRunEASBuild(inputParser stepconf.InputParser) RunEASBuild {
	return RunEASBuild{inputParser: inputParser}
}

// ProcessConfig ...
func (s RunEASBuild) ProcessConfig() (Config, error) {
	var input Input
	err := s.inputParser.Parse(&input)
	if err != nil {
		return Config{}, err
	}
	stepconf.Print(input)

	var platforms []string
	split := strings.Split(input.Platforms, "\n")
	for _, e := range split {
		e = strings.TrimSpace(e)
		if len(e) > 0 {
			platforms = append(platforms, e)
		}
	}
	if len(platforms) == 0 {
		return Config{}, errors.New("no platform specified")
	}

	return Config{
		User:      input.User,
		Password:  input.Password,
		Platforms: platforms,
	}, nil
}

// InstallDeps ...
func (s RunEASBuild) InstallDeps() error {
	return nil
}

// Result ...
type Result struct {
}

// Run ...
func (s RunEASBuild) Run(cfg Config) (Result, error) {
	if err := s.expoClient.Login(cfg.User, cfg.Password); err != nil {
		return Result{}, err
	}

	if err := s.expoClient.Build(cfg.Platforms); err != nil {
		return Result{}, err
	}

	return Result{}, nil
}

// Export ...
func (s RunEASBuild) Export(result Result) error {
	return nil
}
