package fsutil

import (
	"fmt"
	"os"

	"github.com/codeclysm/fs"
)

// Lock wraps ioutil to provide locked read and write
type Lock struct {
	Util
}

func NewLock(fs fs.FS) Lock {
	io := Base{FS: fs}

	return Lock{Util: io}
}

func (l Lock) ReadFile(filename string) ([]byte, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	fmt.Println(stat)

	return l.Util.ReadFile(filename)
}
func (l Lock) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return l.Util.WriteFile(filename, data, perm)
}
