package utils

import "os"

func Modulo(a, b int) int {
	y := a % b
	if y < 0 {
		y = y + b
	}
	return y
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
