package utils

import (
	"slices"
	"testing"
)

func TestBuildArgs_SudoTrue(t *testing.T) {
	args := BuildArgs("apt-get install -y git", true)
	expected := []string{"sudo", "apt-get", "install", "-y", "git"}

	if !slices.Equal(expected, args) {
		t.Errorf("expected %s, got: %s", expected, args)
	}
}

func TestBuildArgs_SudoFalse(t *testing.T) {
	args := BuildArgs("apt-get install -y curl", false)
	expected := []string{"apt-get", "install", "-y", "curl"}

	if !slices.Equal(expected, args) {
		t.Errorf("expected %s, got: %s", expected, args)
	}
}
