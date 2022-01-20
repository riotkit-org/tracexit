package main

import (
	"os"
	"strings"
)

func parseEnvironmentVariable(envAsStr string) (name string, value string) {
	split := strings.SplitN(envAsStr, "=", 2)

	return split[0], split[1]
}

// parseVariableAppending appends text to existing environment variable
// Example input: PATH+BIN=/opt/bin
// Example output: PATH=/bin:/sbin:/usr/bin:/opt/bin
func parseVariableAppending(envAsStr string) string {
	name, valueToAppend := parseEnvironmentVariable(envAsStr)

	// make sure
	if !strings.Contains(name, "+") {
		return envAsStr
	}

	split := strings.SplitN(name, "+", 2)
	realEnvName := split[0]
	newValue := os.Getenv(realEnvName) + ":" + valueToAppend

	return realEnvName + "=" + newValue
}

// ParseEnvironmentVariables returns list of environment variables
// Example input: env:TEST=1 env:BAKUNIN=Mikhail book-reader --from-env
// Example output: {TEST=1, BAKUNIN=Mikhail}
func parseEnvironmentVariables(inputArgs []string) (out []string, remainingArgs []string) {
	for _, arg := range inputArgs {
		if len(arg) < 7 || arg[0:4] != "env:" || !strings.Contains(arg, "=") {
			remainingArgs = append(remainingArgs, arg)
			continue
		}

		envAsStr := arg[4:]

		// support for appending value to variables e.g. PATH+BIN=/opt/bin
		if strings.Contains(envAsStr, "+") {
			envAsStr = parseVariableAppending(envAsStr)
		}

		// side effect: inform OS about the change
		// this allows us to multiple append to same environment variable
		// be aware of that in unit tests!
		realEnvName, newValue := parseEnvironmentVariable(envAsStr)
		_ = os.Setenv(realEnvName, newValue)

		out = append(out, envAsStr)
	}

	return out, remainingArgs
}
