package main

import (
	"github.com/stretchr/testify/assert"
	"os"
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

func TestParseVariableAppending(t *testing.T) {
	_ = os.Setenv("TEST_1", "Kropotkin")
	newVarAsStr := parseVariableAppending("TEST_1+NEW=Bakunin")

	assert.Equal(t, "TEST_1=Kropotkin:Bakunin", newVarAsStr)
}

func TestParseVariableAppending_IgnoresExtraPlus(t *testing.T) {
	_ = os.Setenv("TEST_2", "Kropotkin")
	newVarAsStr := parseVariableAppending("TEST_2+NEW+EXTRAPLUS=Bakunin")

	assert.Equal(t, "TEST_2=Kropotkin:Bakunin", newVarAsStr)
}

func TestParseVariableAppending_DoesNothingIfPlusNotPresent(t *testing.T) {
	_ = os.Setenv("TEST_3", "Kropotkin")
	newVarAsStr := parseVariableAppending("TEST_3=Bakunin")

	assert.Equal(t, "TEST_3=Bakunin", newVarAsStr)
}

func TestParsingEnvironmentVariable(t *testing.T) {
	name, value := parseEnvironmentVariable("HELLO=Anarchism")

	assert.Equal(t, "HELLO", name)
	assert.Equal(t, "Anarchism", value)
}
