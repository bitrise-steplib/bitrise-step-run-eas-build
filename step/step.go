package step

import (
	"github.com/bitrise-io/go-steputils/v2/stepconf"
)

// Input ...
type Input struct {
	User     stepconf.Secret `env:"user,required"`
	Password stepconf.Secret `env:"password,required"`
}

// Config ...
type Config Input

// RunEASBuild ...
type RunEASBuild struct {
	inputParser stepconf.InputParser
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
	stepconf.Print(input) // TODO: log.Infof(stepconf.toString(input))
	return Config(input), nil
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
	return Result{}, nil
}

// Export ...
func (s RunEASBuild) Export(result Result) error {
	return nil
}
