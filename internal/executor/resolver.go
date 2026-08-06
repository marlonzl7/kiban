package executor

import "strings"

func ResolveCommand(cmd string, params map[string]string) string {
	for key, value := range params {
		cmd = strings.ReplaceAll(cmd, "{{"+key+"}}", value)
	}

	return cmd
}
