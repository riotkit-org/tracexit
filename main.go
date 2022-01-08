package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func getProgramArgs() []string {
	return os.Args[2:]
}

func getProcessName() string {
	return os.Args[1]
}

func StartProcess() *exec.Cmd {
	cmd := exec.Command(getProcessName(), getProgramArgs()...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Errorf("Cannot start process: %v", err)
		os.Exit(10)
	}

	return cmd
}

func waitAndRetrieveStatusCode(cmd *exec.Cmd) int {
	code := 0

	if err := cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				code = status.ExitStatus()
			}
		}
	}
	return code
}

func saveExitCode(code int, path string) error {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0644)
	}

	return os.WriteFile(path, []byte(string(rune(code))), 0644)
}

func exit(code int) {
	os.Exit(code)
}

func main() {
	whereToSaveExitCode := os.Getenv("TRACEXIT_EXIT_CODE_PATH")

	if whereToSaveExitCode == "" {
		log.Error("'TRACEXIT_EXIT_CODE_PATH' environment variable needs to be set")
		os.Exit(22)
	}

	if len(os.Args) < 2 {
		log.Error("No command specified.")
		os.Exit(22)
	}

	cmd := StartProcess()
	code := waitAndRetrieveStatusCode(cmd)
	if err := saveExitCode(code, whereToSaveExitCode); err != nil {
		log.Errorf("Cannot write status to file: %v", err)
		exit(5)
	}

	exit(code)
}
