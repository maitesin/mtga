package app_test

import (
	"fmt"
	"testing"

	"github.com/maitesin/mtga/internal/app"
	"github.com/stretchr/testify/require"
)

func TestInvalidCommandError_Error(t *testing.T) {
	t.Parallel()

	cmd1 := app.CreateCardCmd{}
	cmd2 := app.CreateCardCmd{}
	err := app.InvalidCommandError{Received: cmd1, Expected: cmd2}

	require.Contains(t, err.Error(), cmd1.Name())
	require.Contains(t, err.Error(), cmd2.Name())

	require.Equal(t, err.Error(), fmt.Sprintf("invalid command %q received. Expected %q", cmd1.Name(), cmd2.Name()))
}
