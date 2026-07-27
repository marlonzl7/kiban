package detector

import "testing"

func TestDetectDistro_Ubuntu(t *testing.T) {
	info, err := DetectDistro("testdata/ubuntu")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.ID != "ubuntu" {
		t.Errorf("ID: expected %q, got %q", "ubuntu", info.ID)
	}

	if info.IDLike != "debian" {
		t.Errorf("IDLike: expected %q, got %q", "debian", info.IDLike)
	}

	if info.Version != "24.04" {
		t.Errorf("Version: expected %q, got %q", "24.04", info.Version)
	}
}

func TestDetectDistro_LinuxMint(t *testing.T) {
	info, err := DetectDistro("testdata/linuxmint")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.ID != "linuxmint" {
		t.Errorf("ID: expected %q, got %q", "linuxmint", info.ID)
	}

	if info.IDLike != "ubuntu debian" {
		t.Errorf("IDLike: expected %q, got %q", "ubuntu debian", info.IDLike)
	}

	if info.Version != "21" {
		t.Errorf("Version: expected %q, got %q", "21", info.Version)
	}
}

func TestDetectDistro_MissingID(t *testing.T) {
	_, err := DetectDistro("testdata/sem_id")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestDetectDistro_FileNotFound(t *testing.T) {
	_, err := DetectDistro("testdata/nao-existe")

	if err == nil {
		t.Error("expected an error, got nil")
	}
}
