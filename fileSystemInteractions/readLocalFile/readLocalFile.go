package readLocalFile

import (
	"fmt"
	"os"
)

func ReadLocalFile(path string) ([]byte, error) {

	/*
		The ReadFile() method reads the location from its function parameter and returns
		that file's content. That return value is then stored in the data variable. An error is
		returned if the file cannot be read.
		
		ReadFile() is a helper function that calls os.Open() and retrieves an io.Reader.
		The io.ReadAll() function is used to read the entire content of io.Reader.
	*/
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("file read error: %s", err)
	}

	return data, nil
}
