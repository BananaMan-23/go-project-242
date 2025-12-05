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

	if info.IsDir() {
		return fmt.Sprintf("%dB\t%s", info.Size(), path), nil
	}

	value, err := os.ReadDir(path)
	
	if err != nil {
		return "", err
	}
	var total int64

	for _, v := range value {
		if v.Type().IsRegular() {
			info, err := v.Info()

			if err != nil {
				continue
			}
			total += info.Size()
		}
	}

	return fmt.Sprintf("%dB\t%s", total, path), nil
}