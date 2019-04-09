package fs

import (
	"os"
	"path/filepath"
	"time"
)

// Chroot wraps a Filesystem interface modifying paths in order to constrain them to a specific directory
type Chroot struct {
	FS
	Base string
}

func (c Chroot) Abs(name string) (string, error) {
	return filepath.Join(c.Base, name), nil
}

func (c Chroot) Chmod(name string, mode os.FileMode) error {
	name, _ = c.Abs(name)
	return c.FS.Chmod(name, mode)
}

func (c Chroot) Chown(name string, uid, gid int) error {
	name, _ = c.Abs(name)
	return c.FS.Chown(name, uid, gid)
}

func (c Chroot) Chtimes(name string, atime time.Time, mtime time.Time) error {
	name, _ = c.Abs(name)
	return c.FS.Chtimes(name, atime, mtime)
}

func (c Chroot) Create(name string) (*os.File, error) {
	name, _ = c.Abs(name)
	return c.FS.Create(name)
}

func (c Chroot) Lchown(name string, uid, gid int) error {
	name, _ = c.Abs(name)
	return c.FS.Lchown(name, uid, gid)
}

func (c Chroot) Link(oldname, newname string) error {
	oldname, _ = c.Abs(oldname)
	newname, _ = c.Abs(newname)
	return c.FS.Link(oldname, newname)
}

func (c Chroot) Lstat(name string) (os.FileInfo, error) {
	name, _ = c.Abs(name)
	return c.FS.Lstat(name)
}

func (c Chroot) Mkdir(name string, perm os.FileMode) error {
	name, _ = c.Abs(name)
	return c.FS.Mkdir(name, perm)
}

func (c Chroot) MkdirAll(path string, perm os.FileMode) error {
	path, _ = c.Abs(path)
	return c.FS.MkdirAll(path, perm)
}

func (c Chroot) NewFile(fd uintptr, name string) *os.File {
	name, _ = c.Abs(name)
	return c.FS.NewFile(fd, name)
}

func (c Chroot) Open(name string) (*os.File, error) {
	name, _ = c.Abs(name)
	return c.FS.Open(name)
}

func (c Chroot) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	name, _ = c.Abs(name)
	return c.FS.OpenFile(name, flag, perm)
}

func (c Chroot) Readlink(name string) (string, error) {
	name, _ = c.Abs(name)
	return c.FS.Readlink(name)
}

func (c Chroot) Remove(name string) error {
	name, _ = c.Abs(name)
	return c.FS.Remove(name)
}

func (c Chroot) RemoveAll(path string) error {
	path, _ = c.Abs(path)
	return c.FS.RemoveAll(path)
}

func (c Chroot) Rename(oldpath, newpath string) error {
	oldpath, _ = c.Abs(oldpath)
	newpath, _ = c.Abs(newpath)
	return c.FS.Rename(oldpath, newpath)
}

func (c Chroot) Symlink(oldname, newname string) error {
	oldname, _ = c.Abs(oldname)
	newname, _ = c.Abs(newname)
	return c.FS.Symlink(oldname, newname)
}

func (c Chroot) Truncate(name string, size int64) error {
	name, _ = c.Abs(name)
	return c.FS.Truncate(name, size)
}
