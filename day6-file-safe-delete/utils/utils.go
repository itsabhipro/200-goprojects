package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilePath(directory string, filename string) (string, error) {
	dirs, err := os.ReadDir(directory)
	if err != nil {
		return "", err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			path, err := GetFilePath(filepath.Join(directory, dir.Name()), filename)
			if err == nil {
				return path, err
			}
		} else if strings.EqualFold(dir.Name(), filename) || strings.Contains(dir.Name(), filename) {
			return filepath.Join(directory, dir.Name()), nil
		}
	}
	return "", os.ErrNotExist
}
