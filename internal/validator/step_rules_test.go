package validator

import (
	"strings"
	"testing"

	"github.com/marlonzl7/kiban/internal/loader"
)

func TestValidateStepShellSudoExclusivity_NoConflict(t *testing.T) {
	tools := []loader.Tool{
		{
			Name: "git",
			Install: map[string]loader.Install{
				"apt": {
					Steps: []loader.Step{
						{Command: "apt-get install -y git", Sudo: true, Shell: false},
					},
				},
			},
		},
	}

	err := ValidateStepShellSudoExclusivity(tools)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidateStepShellSudoExclusivity_ShellWithoutSudo_OK(t *testing.T) {
	tools := []loader.Tool{
		{
			Name: "docker",
			Install: map[string]loader.Install{
				"apt": {
					Steps: []loader.Step{
						{Command: "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo tee /etc/apt/keyrings/docker.asc > /dev/null", Sudo: false, Shell: true},
					},
				},
			},
		},
	}

	err := ValidateStepShellSudoExclusivity(tools)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidateStepShellSudoExclusivity_ShellAndSudo_Conflict(t *testing.T) {
	tools := []loader.Tool{
		{
			Name: "docker",
			Install: map[string]loader.Install{
				"apt": {
					Steps: []loader.Step{
						{Command: "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | tee /etc/apt/keyrings/docker.asc", Sudo: true, Shell: true},
					},
				},
			},
		},
	}

	err := ValidateStepShellSudoExclusivity(tools)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "docker") {
		t.Errorf("expected error to mention tool name %q, got: %q", "docker", err.Error())
	}
}

func TestValidateStepShellSudoExclusivity_MultiplePackageManagers(t *testing.T) {
	tools := []loader.Tool{
		{
			Name: "docker",
			Install: map[string]loader.Install{
				"apt": {
					Steps: []loader.Step{
						{Command: "curl ... | sudo tee ...", Sudo: false, Shell: true},
					},
				},
				"dnf": {
					Steps: []loader.Step{
						{Command: "curl ... | tee ...", Sudo: true, Shell: true},
					},
				},
			},
		},
	}

	err := ValidateStepShellSudoExclusivity(tools)
	if err == nil {
		t.Fatal("expected an error for the dnf step, got nil")
	}

	if !strings.Contains(err.Error(), "dnf") {
		t.Errorf("expected error to mention package manager %q, got: %q", "dnf", err.Error())
	}
}

func TestValidateStepShellSudoExclusivity_MultipleTools_AggregatesErrors(t *testing.T) {
	tools := []loader.Tool{
		{
			Name: "docker",
			Install: map[string]loader.Install{
				"apt": {Steps: []loader.Step{{Command: "a", Sudo: true, Shell: true}}},
			},
		},
		{
			Name: "rust",
			Install: map[string]loader.Install{
				"apt": {Steps: []loader.Step{{Command: "b", Sudo: true, Shell: true}}},
			},
		},
	}

	err := ValidateStepShellSudoExclusivity(tools)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "docker") || !strings.Contains(err.Error(), "rust") {
		t.Errorf("expected error to mention both tools, got: %q", err.Error())
	}
}
