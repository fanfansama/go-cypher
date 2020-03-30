package miscs

import (
	"os"
	"log"
)

// IsFile : Function to check if a given file path is actually a file
func IsFile(file string) bool {
	res, err := os.Stat(file)

	if err != nil || res.IsDir() {
		return false
	}
	return true
}

// FileSize : Return the file size of a file
func FileSize(file string) (int64, error) {
	res, err := os.Stat(file)
	if err != nil {
		return -1, err
	}
	return res.Size(), nil
}

func DeleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) { return }

	log.Println("==> done deleting file :" + path)
}

func isError(err error) bool {
	if err != nil {
		log.Println(err.Error())
	}

	return (err != nil)
}