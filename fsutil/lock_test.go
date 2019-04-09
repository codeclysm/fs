package fsutil_test

import (
	"io/ioutil"
	"testing"

	"github.com/codeclysm/fs"
	"github.com/codeclysm/fs/fsutil"
)

func TestLock(t *testing.T) {
	tmp, _ := ioutil.TempDir("", "")
	chroot := fs.Chroot{
		FS:   fs.OS{},
		Base: tmp,
	}

	io := fsutil.NewLock(chroot)

	err := io.WriteFile("/file", []byte("hello"), 0644)
	check(t, err)
}
