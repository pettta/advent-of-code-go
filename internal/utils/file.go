package utils

import (
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadFileLines(path string) ([]string, error) {
	lines, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(lines), "\n"), nil
}

func ReadLines(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func ReadCSV(data []byte) []string {
	return strings.Split(string(data), ",")
}