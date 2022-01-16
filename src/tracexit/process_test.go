package tracexit

import (
	"github.com/stretchr/testify/assert"
	"os"
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
