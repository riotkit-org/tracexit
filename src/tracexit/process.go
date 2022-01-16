package tracexit

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func getProcessName(args []string) string {
	return args[0]
}

func StartProcess(processName string, processArgs []string, extraEnv []string) *exec.Cmd {
	cmd := exec.Command(processName, processArgs...)
	if extraEnv != nil {
		cmd.Env = append(os.Environ(), extraEnv...)
	}
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

	return os.WriteFile(path, []byte(fmt.Sprintf("%v", code)), 0644)
}

func exit(code int) {
	os.Exit(code)
}
