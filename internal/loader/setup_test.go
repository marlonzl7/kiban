package loader

import (
	"slices"
	"testing"
)

func TestLoadSetupFile_Success(t *testing.T) {
	sf, err := LoadSetupFile("testdata/setup.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sf.Version != 1 {
		t.Errorf("Version: expected %d, got %d", 1, sf.Version)
	}

	if !slices.Contains(sf.Tools["containers"], "docker") {
		t.Errorf("Name: expected %q, got %q", "docker", sf.Tools["containers"])
	}
}

func TestLoadSetupFile_FileNotFound(t *testing.T) {
	_, err := LoadSetupFile("testdata/nao-existe")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestLoadSetupFile_BadFormat(t *testing.T) {
	_, err := LoadSetupFile("testdata/invalid.yaml")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}
