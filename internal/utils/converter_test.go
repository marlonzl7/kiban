package utils

import (
	"testing"

	"github.com/marlonzl7/kiban/internal/loader"
)

func TestNameToTool_MapsByName(t *testing.T) {
	tools := []loader.Tool{
		{Name: "docker"},
		{Name: "git"},
	}

	result := NameToTool(tools)

	if len(result) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(result))
	}

	if result["docker"].Name != "docker" {
		t.Errorf("expected docker entry, got %+v", result["docker"])
	}

	if result["git"].Name != "git" {
		t.Errorf("expected git entry, got %+v", result["git"])
	}
}

func TestNameToTool_EmptyInput(t *testing.T) {
	result := NameToTool([]loader.Tool{})

	if len(result) != 0 {
		t.Errorf("expected empty map, got %d entries", len(result))
	}
}
