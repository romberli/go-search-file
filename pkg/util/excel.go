package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
)

const (
	defaultXLSXSuffix    = ".xlsx"
	defaultTmpFilePrefix = "~$"
	alphabetNum          = 25
	minAlphabetAscii     = 65
)

func FindExcelFile(path string) ([]string, error) {
	var (
		err   error
		files []string
	)

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, errors.Trace(err)
		}
	}

	err = findExcelFile(path, &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func findExcelFile(path string, files *[]string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		if IsExcel(fileInfo.Name()) {
			*files = append(*files, path)

			return nil
		}

		return errors.Errorf("input path is neither a directory nor a excel file. dir: %s", path)
	}

	paths, err := os.ReadDir(path)
	if err != nil {
		return errors.Trace(err)
	}

	for _, p := range paths {
		pathName := fmt.Sprintf("%s/%s", path, p.Name())

		if p.IsDir() {
			err = findExcelFile(pathName, files)
			if err != nil {
				return err
			}
		} else if IsExcel(p.Name()) {
			*files = append(*files, pathName)
		}
	}

	return nil
}

func IsExcel(file string) bool {
	return strings.HasSuffix(file, defaultXLSXSuffix) && !strings.HasPrefix(file, defaultTmpFilePrefix)
}

func ConvertIntToAlphabet(index int) string {
	p := index / alphabetNum
	q := index % alphabetNum

	column := string(rune(q + minAlphabetAscii))
	if p == constant.ZeroInt {
		return column
	}

	return string(rune(p+minAlphabetAscii-1)) + column
}
