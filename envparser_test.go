package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEnvironmentVariables_Finds_Variables(t *testing.T) {
	out, remainingArgs := parseEnvironmentVariables([]string{"env:A=1", "env:WITH_SPACE=Oopps! This has a space!", "ps", "aux"})

	assert.Equal(t, []string{"A=1", "WITH_SPACE=Oopps! This has a space!"}, out)
	assert.Equal(t, []string{"ps", "aux"}, remainingArgs)
}

func TestParseEnvironmentVariables_WillIgnoreWhenThereIsNoPrefix(t *testing.T) {
	out, remainingArgs := parseEnvironmentVariables([]string{"A=1", "ps", "aux"})

	assert.Equal(t, 0, len(out))
	assert.Equal(t, []string{"A=1", "ps", "aux"}, remainingArgs)
}

func TestParseEnvironmentVariables_CommandMissing(t *testing.T) {
	out, remainingArgs := parseEnvironmentVariables([]string{"env:A=2"})

	assert.Equal(t, []string{"A=2"}, out)
	assert.Equal(t, 0, len(remainingArgs))
}

func TestParseEnvironmentVariables_VariablesWithoutValueShouldBeIgnored(t *testing.T) {
	out, remainingArgs := parseEnvironmentVariables([]string{"env:SOMEVARIABLE"})

	assert.Equal(t, 0, len(out))
	assert.Equal(t, 1, len(remainingArgs))
}
