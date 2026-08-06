package executor

import (
	"strings"
	"testing"

	"github.com/marlonzl7/kiban/internal/loader"
)

func TestVerifyInstallation_Success(t *testing.T) {
	err := VerifyInstallation(loader.Verify{Command: "docker --version", Expect: "Docker version"})
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestVerifyInstallation_Fail(t *testing.T) {
	err := VerifyInstallation(loader.Verify{Command: "docker --versionn", Expect: "Docker version"})
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestVerifyInstallation_VerificationFailed(t *testing.T) {
	err := VerifyInstallation(loader.Verify{Command: "docker --version", Expect: "Java version"})

	if !strings.Contains(err.Error(), "verification failed") {
		t.Errorf("expected verification failed error, got: %v", err)
	}
}
