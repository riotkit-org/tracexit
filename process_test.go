package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestProcessSpawning(t *testing.T) {
	process := StartProcess("echo", []string{"Hello"}, []string{})
	code := waitAndRetrieveStatusCode(process)

	assert.Equal(t, 0, code)
}

func TestProcessExitCodeSaving(t *testing.T) {
	_ = os.Mkdir(".build", 0755)

	process := StartProcess("/bin/true", []string{}, []string{})
	_ = saveExitCode(waitAndRetrieveStatusCode(process), ".build/TestProcessExitCodeSaving")

	actual, _ := os.ReadFile(".build/TestProcessExitCodeSaving")
	assert.Equal(t, "0", string(actual))
}

func TestProcessExitCodeSaving_WhenProcessFails(t *testing.T) {
	_ = os.Mkdir(".build", 0755)

	process := StartProcess("/bin/false", []string{}, []string{})
	_ = saveExitCode(waitAndRetrieveStatusCode(process), ".build/TestProcessExitCodeSaving")

	actual, _ := os.ReadFile(".build/TestProcessExitCodeSaving")
	assert.Equal(t, "1", string(actual))
}

func TestExtraPWDSetting(t *testing.T) {
	_ = os.Mkdir(".build", 0755)
	currentDir, _ := os.Getwd()

	process := StartProcess("/bin/bash",
		[]string{"-c", fmt.Sprintf("pwd > %v/.build/TestExtraPWDSetting", currentDir)},
		[]string{"PWD=/tmp"})
	_ = waitAndRetrieveStatusCode(process)

	actual, _ := os.ReadFile(".build/TestExtraPWDSetting")
	assert.Equal(t, "/tmp", strings.Trim(string(actual), "\n "))
}
