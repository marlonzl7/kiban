package detector

import (
	"fmt"
	"strings"
)

var distroToPackageManager = map[string]string{
	"ubuntu":  "apt",
	"debian":  "apt",
	"fedora":  "dnf",
	"rhel":    "dnf",
	"centos":  "dnf",
	"arch":    "pacman",
	"manjaro": "pacman",
}

func DetectPackageManager(info InfoDistro) (string, error) {
	if pm, exists := distroToPackageManager[info.ID]; exists {
		return pm, nil
	}

	for _, like := range strings.Fields(info.IDLike) {
		if pm, exists := distroToPackageManager[like]; exists {
			return pm, nil
		}
	}

	return "", fmt.Errorf("unsupported distro")
}
