package paths

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// EnsureDirectory 確保傳入資料夾路徑存在, 不存在會主動建立之
func EnsureDirectory(log *logrus.Logger, dir string) error {
	if fi, err := os.Stat(dir); err != nil {
		log.Debugf("Creating %s \n", dir)
		if err = os.MkdirAll(dir, 0777); err != nil {
			return fmt.Errorf("could not create %s: %s", dir, err)
		}
	} else if !fi.IsDir() {
		return fmt.Errorf("%s must be a directory", dir)
	}
	return nil
}

// Exists checks the path exists
func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// IsRegularFile checks the path is a regular file, not a directory
func IsRegularFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.Mode().IsRegular(), nil
}
