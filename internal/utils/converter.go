package utils

import "github.com/marlonzl7/kiban/internal/loader"

func NameToTool(tools []loader.Tool) map[string]loader.Tool {
	toolsMap := make(map[string]loader.Tool)

	for _, tool := range tools {
		toolsMap[tool.Name] = tool
	}

	return toolsMap
}