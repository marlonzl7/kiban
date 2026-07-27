package detector

import "fmt"

var distroToPackageManager = map[string]string{
	"ubuntu":  "apt",
	"debian":  "apt",
	"fedora":  "dnf",
	"rhel":    "dnf",
	"arch":    "pacman",
	"manjaro": "pacman",
}

func DetectPackageManager(info InfoDistro) (string, error) {
	if pm, exists := distroToPackageManager[info.ID]; exists {
		return pm, nil
	}

	if pm, exists := distroToPackageManager[info.IDLike]; exists {
		return pm, nil
	}

	return "", fmt.Errorf("unsupported distro")
}
