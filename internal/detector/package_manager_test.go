package detector

import "testing"

func TestDetectPackageManager_ByID(t *testing.T) {
	info := InfoDistro{ID: "ubuntu"}
	pm, err := DetectPackageManager(info)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pm != "apt" {
		t.Errorf("expected %q, got %q", "apt", pm)
	}
}

func TestDetectPackageManager_FallbackByIDLike(t *testing.T) {
	info := InfoDistro{ID: "linuxmint", IDLike: "ubuntu"}
	pm, err := DetectPackageManager(info)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pm != "apt" {
		t.Errorf("expected %q, got %q", "apt", pm)
	}
}

func TestDetectPackageManager_UnsupportedDistro(t *testing.T) {
	info := InfoDistro{ID: "unknown-distro"}
	_, err := DetectPackageManager(info)

	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestDetectPackageManager_MultipleIDLike(t *testing.T) {
	info := InfoDistro{ID: "almalinux", IDLike: "rhel centos fedora"}
	pm, err := DetectPackageManager(info)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pm != "dnf" {
		t.Errorf("expected %q, got %q", "dnf", pm)
	}
}

func TestDetectPackageManager_IDLikeFirstUnknownSecondKnown(t *testing.T) {
	info := InfoDistro{ID: "foo", IDLike: "unknown ubuntu debian"}
	pm, err := DetectPackageManager(info)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pm != "apt" {
		t.Errorf("expected %q, got %q", "apt", pm)
	}
}