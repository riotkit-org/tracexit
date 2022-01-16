package tracexit

import "strings"

// ParseEnvironmentVariables returns list of environment variables
// Example input: env:TEST=1 env:BAKUNIN=Mikhail book-reader --from-env
// Example output: {TEST=1, BAKUNIN=Mikhail}
func parseEnvironmentVariables(inputArgs []string) (out []string, remainingArgs []string) {
	for _, arg := range inputArgs {
		if len(arg) < 7 || arg[0:4] != "env:" || !strings.Contains(arg, "=") {
			remainingArgs = append(remainingArgs, arg)
			continue
		}

		out = append(out, arg[4:])
	}

	return out, remainingArgs
}
