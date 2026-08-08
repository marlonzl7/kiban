package utils

import (
	"slices"
	"strings"
)

func BuildArgs(command string, sudo bool) []string {
	args := strings.Split(command, " ")

	if sudo {
		args = slices.Insert(args, 0, "sudo")
	}

	return args
}
