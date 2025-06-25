// Package utils contains utility logic
package utils

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// FileExists checks if the passed file exists
func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FileIsEmpty checks if the file is empty
func FileIsEmpty(filename string) bool {
	info, _ := os.Stat(filename)
	if info.Size() == 0 {
		return true
	}
	return false
}

// GetLineCount returns the number of lines of the passed file
func GetLineCount(fileName string) (int, error) {
	if FileExists(fileName) == false {
		return 0, nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer CloseFileAndHandleError(file, &err)

	csvReader := csv.NewReader(file)
	rowCount := 0
	for {
		_, err = csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		rowCount++
	}

	return rowCount, nil
}

// CloseFileAndHandleError closes the file and handles a potential error
func CloseFileAndHandleError(file *os.File, err *error) {
	closeErr := file.Close()
	if closeErr != nil {
		if *err == nil { // Only overwrite if no prior error
			*err = closeErr
		}
	}
}

// ClearFile clears the passed file
func ClearFile(fileName string) error {
	return os.Truncate(fileName, 0)
}

// ToBool converts a string to a boolean value
func ToBool(info string) bool {
	aBool, _ := strconv.ParseBool(info)
	return aBool
}

// IntToString converts an int to a string value
func IntToString(info int) string {
	aString := strconv.Itoa(info)
	return aString
}
