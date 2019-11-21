package releasestructure

import (
	"errors"
	"fmt"
	"os"
)

//IsAvailable does check if the directory is available
func IsAvailable(pathToDirectory string) (bool, error) {
	_, err := os.Stat(pathToDirectory)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		errorString := fmt.Sprintf("directory %s does not exist or is on the wrong place", pathToDirectory)
		return false, errors.New(errorString)
	}

	if !os.IsNotExist(err) {
		// path/to/whatever exists
		return true, nil
	}

	return false, err

}
func check(pathToDirectory string) bool {
	if boolean, err := IsAvailable(pathToDirectory); boolean == true && err == nil {
		return true
	}
	return false
}

func main() {
	if check("../../bosh-release_test/packages") == true {
		fmt.Printf("directory packages exists\n")
	} else {
		fmt.Printf("directory packages doesn't exists\n")
	}

	if check("../../bosh-release_test/src") == true {
		fmt.Printf("directory jobs exists\n")
	} else {
		fmt.Printf("directory jobs doesn't exists\n")
	}

	if check("../../bosh-release_test/config") == true {
		fmt.Printf("directory config exists \n")
	} else {
		fmt.Printf("directory congif doesn't exists\n")
	}

}
