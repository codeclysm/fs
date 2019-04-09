package fs

import (
	"fmt"
	"os"
)

// Lock wraps ioutil to provide locked read and write
type Lock struct {
	Ioutil
}

func NewLock(fs FS) Lock {
	io := Base{FS: fs}

	return Lock{Ioutil: io}
}

func (l Lock) ReadFile(filename string) ([]byte, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	fmt.Println(stat)

	return l.Ioutil.ReadFile(filename)
}
func (l Lock) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return l.Ioutil.WriteFile(filename, data, perm)
}
