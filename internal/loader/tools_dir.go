package loader

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func LoadToolsFromDir(dir string) ([]Tool, error) {
	var tools []Tool
	var errs []error

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(d.Name(), ".yaml") {
			return nil
		}

		tool, err := LoadToolFile(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", path, err))
			return nil
		}

		tools = append(tools, tool)

		return nil
	})

	if err != nil {
		errs = append(errs, fmt.Errorf("failed to load tools: %w", err))
	}

	return tools, errors.Join(errs...)
}
