package ufiles

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/congta/ugo-commons/commons-logging/ulogs"
	"github.com/congta/ugo-commons/commons-u/ucommons"
	"github.com/duke-git/lancet/v2/fileutil"
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

func IsDir(fileName string) (bool, error) {
	s, err := os.Stat(fileName)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

func IsNormalFile(fileName string) (bool, error) {
	s, err := os.Stat(fileName)
	if err != nil {
		return false, err
	}
	return !s.IsDir(), nil
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

func MustListFiles(dir string) []string {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	ucommons.AssertNonErr(err)
	return files
}

func MustReadLines(fp string) []string {
	lines, err := fileutil.ReadFileByLine(fp)
	ucommons.AssertNonErr(err)
	return lines
}

func MustRelative(base, target string) string {
	rel, err := filepath.Rel(base, target)
	ucommons.AssertNonErr(err)
	return rel
}

func MustRead(fp string) []byte {
	data, err := os.ReadFile(fp)
	ucommons.AssertNonErr(err)
	return data
}

func MustOverwriteLinesToFile(path string, lines []string) {
	_ = fileutil.ClearFile(path)
	_ = fileutil.CreateDir(filepath.Dir(path))
	err := fileutil.WriteStringToFile(path, strings.Join(lines, "\n"), false)
	ucommons.AssertNonErr(err)
}

func MustOverwriteToFile(path string, data []byte) {
	_ = fileutil.ClearFile(path)
	_ = fileutil.CreateDir(filepath.Dir(path))
	err := fileutil.WriteBytesToFile(path, data)
	ucommons.AssertNonErr(err)
}

func MustAppendLinesToFile(path string, lines []string) {
	_ = fileutil.CreateDir(filepath.Dir(path))
	err := fileutil.WriteStringToFile(path, strings.Join(lines, "\n"), true)
	ucommons.AssertNonErr(err)
}
