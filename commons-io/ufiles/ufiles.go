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

func CloseQuietly(file *os.File) {
	if err := file.Close(); err != nil {
		ulogs.Warn("close file error, %v", err)
	}
}
