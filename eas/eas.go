package eas

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
)

type Client struct {
	commandFactory command.Factory
	logger         log.Logger
	token          stepconf.Secret
	workDir        string
}

func NewClient(commandFactory command.Factory, logger log.Logger, token stepconf.Secret, workDir string) Client {
	return Client{
		commandFactory: commandFactory,
		logger:         logger,
		token:          token,
		workDir:        workDir,
	}
}

func (c Client) Build(platform string) error {
	cmd := c.commandFactory.Create("npx", []string{"eas-cli", "build", "--platform", platform, "--non-interactive"}, &command.Opts{
		Env:    []string{"EXPO_TOKEN=" + string(c.token)},
		Dir:    c.workDir,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	c.logger.Donef("$ %s", cmd.PrintableCommandArgs())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %w", cmd.PrintableCommandArgs(), err)
	}
	return nil
}
