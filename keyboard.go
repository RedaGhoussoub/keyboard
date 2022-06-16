// Package keyboard a helper package to retrieve values from standard console input

package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat retrieves a floating point number from console input or an error when input could not be read
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}

	return number, nil
}

// GetInt retrieves an integer from console input or an error when input could not be read
func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)

	if err != nil {
		return 0, err
	}

	return number, nil
}
