package loader

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func LoadToolsFromDir(dir string) ([]Tool, error) {
	var tools []Tool
	var errs []error

	entries, err := os.ReadDir(dir)
	if err != nil {
		return tools, fmt.Errorf("failed to read tools dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		tool, err := LoadToolFile(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", path, err))
			continue
		}

		tools = append(tools, tool)
	}

	return tools, errors.Join(errs...)
}
