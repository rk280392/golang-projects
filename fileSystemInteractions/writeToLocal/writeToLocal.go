package writeToLocal

import "os"

func WriteToLocal(path string, data string) error {
	byteData := []byte(data)
	if err := os.WriteFile(path, byteData, 0644); err != nil {
		return err
	}
	return nil
}
