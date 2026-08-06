package executor

import "testing"

func TestResolveCommand_Full(t *testing.T) {
	params := map[string]string{
		"version": "21",
		"arch":    "x86_64",
	}

	cmd := ResolveCommand("apt-get install -y open-jdk-{{version}}-jdk:{{arch}}", params)

	expected := "apt-get install -y open-jdk-21-jdk:x86_64"

	if cmd != expected {
		t.Errorf("expected '%s', got: %s", expected, cmd)
	}
}

func TestResolveCommand_MissingArch(t *testing.T) {
	params := map[string]string{
		"version": "21",
	}

	cmd := ResolveCommand("apt-get install -y open-jdk-{{version}}-jdk", params)

	expected := "apt-get install -y open-jdk-21-jdk"

	if cmd != expected {
		t.Errorf("expected '%s', got: %s", expected, cmd)
	}
}
