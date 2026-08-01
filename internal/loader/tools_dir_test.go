package loader

import (
	"testing"
)

func TestLoadToolsFromDir_AllValid(t *testing.T) {
	tools, err := LoadToolsFromDir("testdata/tools_valid")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tools) != 2 {
		t.Fatalf("expected 2 tools, got %d", len(tools))
	}
}

func TestLoadToolsFromDir_NestedFolders(t *testing.T) {
	tools, err := LoadToolsFromDir("testdata/tools_nested")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tools) != 2 {
		t.Fatalf("expected 2 tools, got %d", len(tools))
	}
}

func TestLoadToolsFromDir_DirNotFound(t *testing.T) {
	_, err := LoadToolsFromDir("testdata/does-not-exist")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestLoadToolsFromDir_MixedValidAndInvalid(t *testing.T) {
	tools, err := LoadToolsFromDir("testdata/tools_mixed")

	if err == nil {
		t.Error("expected an error due to broken.yaml, got nil")
	}

	if len(tools) != 2 {
		t.Fatalf("expected 2 valid tools loaded despite the error, got %d", len(tools))
	}
}
