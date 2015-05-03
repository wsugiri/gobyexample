package file

import (
	"os"
)

func Mkdir(path string) error {
	return os.Mkdir(path, 0777)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, 0777)
}

func CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}
