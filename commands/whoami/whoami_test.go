package whoami_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/yakuter/gol/commands/whoami"
)

func TestEcho(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			whoami.Command(),
		},
	}

	testArgs := []string{execName, "whoami"}
	require.NoError(t, app.Run(testArgs))
}
