package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Printf("Usage: %v env:SOME=env_value my-application --option1 --option2 -a -b -c\n", os.Args[0])
		fmt.Println("")
		fmt.Println("Functionality")
		fmt.Println("=============")
		fmt.Println("- Receive exit code as text file when process ends: Set TRACEXIT_EXIT_CODE_PATH to a path to have exit code written to file at selected path")
		fmt.Println("- Set any environment variable before process name, so it will be passed to inside the process only. Helpful in environments, where you cannot adjust shell settings and can run only single, encapsulated process e.g. ['my-process', '--my-arg']")
		fmt.Println("")
		os.Exit(1)
	}

	environmentToApply, remainingArgs := parseEnvironmentVariables(os.Args[1:])

	if len(remainingArgs) < 1 {
		log.Error("No command specified.")
		os.Exit(22)
	}

	cmd := StartProcess(getProcessName(remainingArgs), remainingArgs[1:], environmentToApply)
	code := waitAndRetrieveStatusCode(cmd)

	whereToSaveExitCode := os.Getenv("TRACEXIT_EXIT_CODE_PATH")
	println("!!!", whereToSaveExitCode)
	if whereToSaveExitCode != "" {
		if err := saveExitCode(code, whereToSaveExitCode); err != nil {
			log.Errorf("Cannot write status to file: %v", err)
			exit(5)
		}
	}

	exit(code)
}
