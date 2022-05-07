package ufiles

import (
	"bufio"
	"congta.com/ugo-commons/commons-logging/ulogs"
	"os"
)

func ReadLinesTry0(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseQuietly(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func ReadLinesTry(fileName string) []string {
	lines, err := ReadLines(fileName)
	if err != nil {
		ulogs.Warn("read lines from file error, %v", err)
	}
	return lines
}

func WriteLines(fileName string, lines []string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer CloseQuietly(file)

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	// file.Close don't trigger flush
	return writer.Flush()
}

// Exists return (false, err) means it's not sure due to some errors
func Exists(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if err == nil {
		return true, nil
	}
	// error may be non-exist or others
	if os.IsExist(err) {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CloseQuietly(file *os.File) {
	if err := file.Close(); err != nil {
		ulogs.Warn("close file error, %v", err)
	}
}
