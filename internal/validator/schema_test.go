package validator

import (
	"strings"
	"testing"

	"github.com/marlonzl7/kiban/internal/loader"
)

func TestValidate_AllSupported(t *testing.T) {
	sf, err := loader.LoadSetupFile("testdata/setup_valid.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	catalog := []loader.Tool{
		{Name: "docker"},
		{Name: "git"},
	}

	err = Validate(sf, catalog)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidate_UnsupportedTool(t *testing.T) {
	sf, err := loader.LoadSetupFile("testdata/setup_unsupported.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	catalog := []loader.Tool{
		{Name: "docker"},
	}

	err = Validate(sf, catalog)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "unknown-tool") {
		t.Errorf("Error message: expected %q, got %q", "tool unknown-tool is not supported", err.Error())
	}
}

func TestValidate_EmptyToolName(t *testing.T) {
	sf, err := loader.LoadSetupFile("testdata/setup_empty_name.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	catalog := []loader.Tool{
		{Name: "docker"},
	}

	err = Validate(sf, catalog)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "missing tool name") {
		t.Errorf("Error message: expected %q, got %q", "invalid tool entry @21: missing tool name", err.Error())
	}
}
