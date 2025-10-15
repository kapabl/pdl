package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type VerbosePrinter struct {
	enabled bool
}

func NewVerbosePrinter(enabled bool) VerbosePrinter {
	var result VerbosePrinter
	result.enabled = enabled
	return result
}

func (printer VerbosePrinter) Println(message string) {
	if printer.enabled {
		fmt.Println(message)
	}
}

func CleanDir(target string) error {
	if target == "" {
		return errors.New("output directory not specified")
	}
	info, statErr := os.Stat(target)
	if statErr == nil && !info.IsDir() {
		return errors.New("output path exists and is not a directory")
	}
	if statErr == nil {
		oldName := target + "_old"
		suffix := 1
		for {
			renErr := os.Rename(target, oldName)
			if renErr == nil {
				break
			}
			if !os.IsExist(renErr) {
				break
			}
			oldName = target + "_old" + strconv.Itoa(suffix)
			suffix++
		}
		os.RemoveAll(oldName)
	}
	return os.MkdirAll(target, 0o755)
}

func EnsureDir(target string) error {
	if target == "" {
		return errors.New("directory not specified")
	}
	return os.MkdirAll(target, 0o755)
}

func ResolvePath(baseDir string, value string) string {
	if filepath.IsAbs(value) {
		return value
	}
	result := filepath.Join(baseDir, value)
	return result
}
