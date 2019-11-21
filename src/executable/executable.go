package executable

import (
	"errors"
	"fmt"
	"os"
)

func IsExecutable(filename string) (bool, error) {
	fi, err := os.Lstat(filename)
	if err != nil {

		errorString := fmt.Sprintf("file %s does not exist", filename)
		return false, errors.New(errorString)
	}

	if !fi.Mode().IsRegular() {
		errorString := fmt.Sprintf("file %s is not a regular file", filename)
		return false, errors.New(errorString)
	}

	if fi.Mode().Perm()&0111 != 0 {
		return true, nil
	}
	return false, nil
}
