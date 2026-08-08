package loader

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

func LoadToolsFromDir(fsys fs.FS, root string) ([]Tool, error) {
	var tools []Tool
	var errs []error

	err := fs.WalkDir(fsys, root, func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(d.Name(), ".yaml") {
			return nil
		}

		tool, err := LoadToolFile(fsys, name)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
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
