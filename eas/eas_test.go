package eas

import (
	"os"
	"testing"

	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/stretchr/testify/require"
)

func TestClient_Build(t *testing.T) {
	cmd := new(MockCommand)
	cmd.On("PrintableCommandArgs").Return("").Once()
	cmd.On("Run").Return(nil).Once()

	commandFactory := new(MockFactory)
	commandFactory.On("Create", "npx", []string{"eas-cli", "build", "--platform", "all", "--non-interactive"}, &command.Opts{
		Env:    []string{"EXPO_TOKEN=token"},
		Dir:    "$HOME",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}).Return(cmd).Once()

	client := NewClient(commandFactory, log.NewLogger(), "token", "$HOME")
	err := client.Build("all")
	require.NoError(t, err)

	commandFactory.AssertExpectations(t)
	cmd.AssertExpectations(t)
}
