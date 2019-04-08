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

func (c Chroot) join(path string) string {
	return filepath.Join(c.Base, path)
}

func (c Chroot) Chmod(name string, mode os.FileMode) error {
	return c.FS.Chmod(c.join(name), mode)
}

func (c Chroot) Chown(name string, uid, gid int) error {
	return c.FS.Chown(c.join(name), uid, gid)
}

func (c Chroot) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return c.FS.Chtimes(c.join(name), atime, mtime)
}

func (c Chroot) Create(name string) (*os.File, error) {
	return c.FS.Create(c.join(name))
}

func (c Chroot) Lchown(name string, uid, gid int) error {
	return c.FS.Lchown(c.join(name), uid, gid)
}

func (c Chroot) Link(oldname, newname string) error {
	return c.FS.Link(c.join(oldname), c.join(newname))
}

func (c Chroot) Mkdir(name string, perm os.FileMode) error {
	return c.FS.Mkdir(c.join(name), perm)
}

func (c Chroot) MkdirAll(path string, perm os.FileMode) error {
	return c.FS.MkdirAll(c.join(path), perm)
}

func (c Chroot) NewFile(fd uintptr, name string) *os.File {
	return c.FS.NewFile(fd, c.join(name))
}

func (c Chroot) Open(name string) (*os.File, error) {
	return c.FS.Open(c.join(name))
}

func (c Chroot) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return c.FS.OpenFile(c.join(name), flag, perm)
}

func (c Chroot) Readlink(name string) (string, error) {
	return c.FS.Readlink(c.join(name))
}

func (c Chroot) Remove(name string) error {
	return c.FS.Remove(c.join(name))
}

func (c Chroot) RemoveAll(path string) error {
	return c.FS.RemoveAll(c.join(path))
}

func (c Chroot) Rename(oldpath, newpath string) error {
	return c.FS.Rename(c.join(oldpath), c.join(newpath))
}

func (c Chroot) Symlink(oldname, newname string) error {
	return c.FS.Symlink(c.join(oldname), c.join(newname))
}

func (c Chroot) Truncate(name string, size int64) error {
	return c.FS.Truncate(c.join(name), size)
}
