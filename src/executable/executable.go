package executable
import (
	
	"os"
	"fmt"
	"errors"
)
func IsExecutable(filename string) (bool, error) {	
	 fi, err := os.Lstat(filename)
	 if err != nil {
		
		errorString := fmt.Sprintf("file %s does not exist", filename)
		return false, errors.New(errorString)
	 }

  if fi.Mode().Perm() == 0711 {
			return true, nil
  		}
		 return false, nil
}


