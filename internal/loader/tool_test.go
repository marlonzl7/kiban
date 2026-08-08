package loader

import (
	"testing"
)

func TestLoadToolFile_Success(t *testing.T) {
	tool, err := LoadToolFile(toolsFS, "docker.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tool.Name != "docker" {
		t.Errorf("Name: expected %q, got %q", "docker", tool.Name)
	}

	steps := tool.Install["apt"].Steps

	if len(steps) != 2 {
		t.Fatalf("Steps: expected 2, got %d", len(steps))
	}

	if steps[0].Command != "apt-get update" {
		t.Errorf("Steps[0].Command: expected %q, got %q", "apt-get update", steps[0].Command)
	}

	if !steps[0].Sudo {
		t.Error("Steps[0].Sudo: expected true, got false")
	}
}

func TestLoadToolFile_FileNotFound(t *testing.T) {
	_, err := LoadToolFile(toolsFS, "nao-existe")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestLoadToolFile_BadFormat(t *testing.T) {
	_, err := LoadToolFile(toolsFS, "invalid.yaml")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}
