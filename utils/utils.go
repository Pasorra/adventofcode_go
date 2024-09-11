package utils

import (
	"log"
	"os"
	"strings"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// For Read to work, go into the dir of the file to run the go run . command
func Read() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	lines := strings.Split(string(data), "\n")

	for i := range lines {
		lines[i] = strings.Trim(lines[i], "\n\t\r")
	}

	return lines
}

func Contains[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IsDigit(char rune) bool {
	return 48 <= char && char <= 57
}

func Max[T Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
