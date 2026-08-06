package executor

import (
	"os"
	"testing"

	"github.com/marlonzl7/kiban/internal/loader"
)

func TestRunStep_Success(t *testing.T) {
	err := RunStep(loader.Step{Command: "echo hello", Sudo: false})
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestRunStep_Fail(t *testing.T) {
	err := RunStep(loader.Step{Command: "jjjj", Sudo: false})
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestRunSteps_Sucess(t *testing.T) {
	var steps []loader.Step
	steps = append(steps, loader.Step{Command: "echo hello", Sudo: false})
	steps = append(steps, loader.Step{Command: "echo world", Sudo: false})

	err := RunSteps(steps, "", "")
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestRunSteps_Fail(t *testing.T) {
	var steps []loader.Step
	steps = append(steps, loader.Step{Command: "false", Sudo: false})
	steps = append(steps, loader.Step{Command: "touch /tmp/fail", Sudo: false})

	err := RunSteps(steps, "", "")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	if _, err := os.ReadFile("/tmp/fail"); err == nil {
		os.Remove("/tmp/fail")
		t.Errorf("expected the second step not to execute")
	}
}
