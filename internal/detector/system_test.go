package detector

import "testing"

func TestDetectArchitecture_ReturnsNonEmpty(t *testing.T) {
	arch, err := DetectArchitecture()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if arch == "" {
		t.Error("expected a non-empty architecture")
	}
}

func TestIsSudoAvailable_DoesNotPanic(t *testing.T) {
	// this only ensures the function runs without panicking and returns a bool;
	// it does not assert which value is "correct" since that depends on the environment
	_ = IsSudoAvailable()
}
