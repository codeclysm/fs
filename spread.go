package fs

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Spread wraps a filesystem and spreads files in directories accordin to their name
// in order not to have a single folder with millions of subdirectories.
// For example the filename test/banana will be spread in te/st/test/banana
// Make sure to have filenames longer than 4 characters, and to have filenames who can
// ensure a balanced distribution (eg with a random string in front)
// All functions will panic if the file is shorter than 4 characters
type Spread struct {
	FS
}

func (s Spread) spread(path string) string {
	if len(path) < 4 {
		panic("cannot spread a path shorter than 4 characters")
	}
	spreaded := strings.Replace(path, string(filepath.Separator), "", -1)

	// Ensure the spread directories exists
	err := s.FS.MkdirAll(filepath.Join(spreaded[0:2], spreaded[2:4]), 0755)
	if err != nil {
		panic("cannot create spread directories")
	}

	return filepath.Join(spreaded[0:2], spreaded[2:4], path)
}

func (s Spread) Chmod(name string, mode os.FileMode) error {
	return s.FS.Chmod(s.spread(name), mode)
}

func (s Spread) Chown(name string, uid, gid int) error {
	return s.FS.Chown(s.spread(name), uid, gid)
}

func (s Spread) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return s.FS.Chtimes(s.spread(name), atime, mtime)
}

func (s Spread) Create(name string) (*os.File, error) {
	return s.FS.Create(s.spread(name))
}

func (s Spread) Lchown(name string, uid, gid int) error {
	return s.FS.Lchown(s.spread(name), uid, gid)
}

func (s Spread) Link(oldname, newname string) error {
	return s.FS.Link(s.spread(oldname), s.spread(newname))
}

func (s Spread) Mkdir(name string, perm os.FileMode) error {
	return s.FS.Mkdir(s.spread(name), perm)
}

func (s Spread) MkdirAll(path string, perm os.FileMode) error {
	return s.FS.MkdirAll(s.spread(path), perm)
}

func (s Spread) NewFile(fd uintptr, name string) *os.File {
	return s.FS.NewFile(fd, s.spread(name))
}

func (s Spread) Open(name string) (*os.File, error) {
	return s.FS.Open(s.spread(name))
}

func (s Spread) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return s.FS.OpenFile(s.spread(name), flag, perm)
}

func (s Spread) Readlink(name string) (string, error) {
	return s.FS.Readlink(s.spread(name))
}

func (s Spread) Remove(name string) error {
	return s.FS.Remove(s.spread(name))
}

func (s Spread) RemoveAll(path string) error {
	return s.FS.RemoveAll(s.spread(path))
}

func (s Spread) Rename(oldpath, newpath string) error {
	return s.FS.Rename(s.spread(oldpath), s.spread(newpath))
}

func (s Spread) Symlink(oldname, newname string) error {
	return s.FS.Symlink(s.spread(oldname), s.spread(newname))
}

func (s Spread) Truncate(name string, size int64) error {
	return s.FS.Truncate(s.spread(name), size)
}
