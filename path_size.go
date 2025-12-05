package pathSize

import (
	"fmt"
	"os"
)

func GetSize(path string) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmt.Sprintf("%dB\t%s", info.Size(), path), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var total int64
	for _, e := range entries {
		if !e.Type().IsRegular() {
			continue
		}
		fi, err := e.Info()
		if err != nil {
			continue
		}
		total += fi.Size()
	}

	return fmt.Sprintf("%dB\t%s", total, path), nil
}
