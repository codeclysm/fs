package fs

import (
	"errors"
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

func (s Spread) Abs(path string) (string, error) {
	path, err := s.abs(path)
	if err != nil {
		return "", err
	}
	return s.FS.Abs(path)
}

func (s Spread) abs(path string) (string, error) {
	if len(path) < 4 {
		return "", errors.New("cannot spread a path shorter than 4 characters")
	}
	spreaded := strings.Replace(path, string(filepath.Separator), "", -1)

	// Ensure the spread directories exists
	err := s.FS.MkdirAll(filepath.Join(spreaded[0:2], spreaded[2:4]), 0755)
	if err != nil {
		return "", errors.New("cannot create spread directories")
	}

	return filepath.Join(spreaded[0:2], spreaded[2:4], path), nil
}

func (s Spread) Chmod(name string, mode os.FileMode) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Chmod(name, mode)
}

func (s Spread) Chown(name string, uid, gid int) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Chown(name, uid, gid)
}

func (s Spread) Chtimes(name string, atime time.Time, mtime time.Time) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Chtimes(name, atime, mtime)
}

func (s Spread) Create(name string) (*os.File, error) {
	name, err := s.abs(name)
	if err != nil {
		return nil, err
	}
	return s.FS.Create(name)
}

func (s Spread) Lchown(name string, uid, gid int) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Lchown(name, uid, gid)
}

func (s Spread) Link(oldname, newname string) error {
	oldname, err := s.abs(oldname)
	if err != nil {
		return err
	}
	newname, err = s.abs(newname)
	if err != nil {
		return err
	}
	return s.FS.Link(oldname, newname)
}

func (s Spread) Lstat(name string) (os.FileInfo, error) {
	name, err := s.abs(name)
	if err != nil {
		return nil, err
	}
	return s.FS.Lstat(name)
}

func (s Spread) Mkdir(name string, perm os.FileMode) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Mkdir(name, perm)
}

func (s Spread) MkdirAll(path string, perm os.FileMode) error {
	path, err := s.abs(path)
	if err != nil {
		return err
	}
	return s.FS.MkdirAll(path, perm)
}

func (s Spread) NewFile(fd uintptr, name string) *os.File {
	name, err := s.abs(name)
	if err != nil {
		return nil
	}
	return s.FS.NewFile(fd, name)
}

func (s Spread) Open(name string) (*os.File, error) {
	name, err := s.abs(name)
	if err != nil {
		return nil, err
	}
	return s.FS.Open(name)
}

func (s Spread) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	name, err := s.abs(name)
	if err != nil {
		return nil, err
	}
	return s.FS.OpenFile(name, flag, perm)
}

func (s Spread) Readlink(name string) (string, error) {
	name, err := s.abs(name)
	if err != nil {
		return "", err
	}
	return s.FS.Readlink(name)
}

func (s Spread) Remove(name string) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Remove(name)
}

func (s Spread) RemoveAll(path string) error {
	path, err := s.abs(path)
	if err != nil {
		return err
	}
	return s.FS.RemoveAll(path)
}

func (s Spread) Rename(oldpath, newpath string) error {
	oldpath, err := s.abs(oldpath)
	if err != nil {
		return err
	}
	newpath, err = s.abs(newpath)
	if err != nil {
		return err
	}
	return s.FS.Rename(oldpath, newpath)
}

func (s Spread) Symlink(oldname, newname string) error {
	oldname, err := s.abs(oldname)
	if err != nil {
		return err
	}
	newname, err = s.abs(newname)
	if err != nil {
		return err
	}
	return s.FS.Symlink(oldname, newname)
}

func (s Spread) Truncate(name string, size int64) error {
	name, err := s.abs(name)
	if err != nil {
		return err
	}
	return s.FS.Truncate(name, size)
}
