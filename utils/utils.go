package utils

import (
	"os"
	"strings"
)

func ReadFile() []string {
	fileBytes, _ := os.ReadFile("input.txt")

	return strings.Split(string(fileBytes), "\n")
}
