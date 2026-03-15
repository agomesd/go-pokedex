package utils

import (
	"strings"
)

func CleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	lowered := strings.ToLower((trimmed))
	split := strings.Split(lowered, " ")
	return split
}
