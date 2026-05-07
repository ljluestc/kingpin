package kingpin

import (
	stdflag "flag"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandLineArgsPrefersRemainingStdlibArgsWhenConsumed(t *testing.T) {
	origArgs := os.Args
	origFlagSet := stdflag.CommandLine
	defer func() {
		os.Args = origArgs
		stdflag.CommandLine = origFlagSet
	}()

	flagSet := stdflag.NewFlagSet("test", stdflag.ContinueOnError)
	flagSet.SetOutput(io.Discard)
	flagSet.Bool("verbose", false, "")
	stdflag.CommandLine = flagSet

	os.Args = []string{"cmd", "--verbose", "serve"}
	err := stdflag.CommandLine.Parse(os.Args[1:])
	assert.NoError(t, err)

	assert.Equal(t, []string{"serve"}, commandLineArgs())
}

func TestCommandLineArgsFallsBackToOSArgsWhenStdlibDidNotConsume(t *testing.T) {
	origArgs := os.Args
	origFlagSet := stdflag.CommandLine
	defer func() {
		os.Args = origArgs
		stdflag.CommandLine = origFlagSet
	}()

	flagSet := stdflag.NewFlagSet("test", stdflag.ContinueOnError)
	flagSet.SetOutput(io.Discard)
	flagSet.Bool("verbose", false, "")
	stdflag.CommandLine = flagSet

	os.Args = []string{"cmd", "serve"}
	err := stdflag.CommandLine.Parse(os.Args[1:])
	assert.NoError(t, err)

	assert.Equal(t, []string{"serve"}, commandLineArgs())
}

func TestCommandLineArgsUsesOSArgsWhenStdlibNotParsed(t *testing.T) {
	origArgs := os.Args
	origFlagSet := stdflag.CommandLine
	defer func() {
		os.Args = origArgs
		stdflag.CommandLine = origFlagSet
	}()

	flagSet := stdflag.NewFlagSet("test", stdflag.ContinueOnError)
	flagSet.SetOutput(io.Discard)
	flagSet.Bool("verbose", false, "")
	stdflag.CommandLine = flagSet

	os.Args = []string{"cmd", "--verbose", "serve"}

	assert.Equal(t, []string{"--verbose", "serve"}, commandLineArgs())
}
