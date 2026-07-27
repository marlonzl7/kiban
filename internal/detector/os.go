package detector

import (
	"fmt"
	"os"
	"strings"
)

type InfoDistro struct {
	ID      string
	IDLike  string
	Version string
}

func DetectDistro(path string) (InfoDistro, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return InfoDistro{}, fmt.Errorf("failed to detect distro: %w", err)
	}

	var idRelease, idLikeRelease, versionRelease string

	lines := strings.Split(string(data), "\n")

	for i := range lines {
		if strings.HasPrefix(lines[i], "ID=") {
			idRelease = normalize(lines[i], "ID=")
		}

		if strings.HasPrefix(lines[i], "ID_LIKE=") {
			idLikeRelease = normalize(lines[i], "ID_LIKE=")
		}

		if strings.HasPrefix(lines[i], "VERSION_ID=") {
			versionRelease = normalize(lines[i], "VERSION_ID=")
		}
	}

	if idRelease == "" {
		return InfoDistro{}, fmt.Errorf("could not detect distro id")
	}

	return InfoDistro{ID: idRelease, IDLike: idLikeRelease, Version: versionRelease}, nil
}

func normalize(s string, prefix string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, prefix)
	s = strings.Trim(s, "\"")
	return s
}
