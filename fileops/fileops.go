package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	value, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
	fmt.Println("File updated! New value:", valueText)
}
