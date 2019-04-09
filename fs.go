// Package fs is a collection of composable wrappers around the filesystem, useful for when a function accepts
// an interface rather than writing directly on the filesystem
package fs

import (
	"os"
	"time"
)

type FS interface {
	Chmod(name string, mode os.FileMode) error
	Chown(name string, uid, gid int) error
	Chtimes(name string, atime time.Time, mtime time.Time) error
	Create(name string) (*os.File, error)
	Lchown(name string, uid, gid int) error
	Link(oldname, newname string) error
	Lstat(name string) (os.FileInfo, error)
	Mkdir(name string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
	NewFile(fd uintptr, name string) *os.File
	Open(name string) (*os.File, error)
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	Readlink(name string) (string, error)
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error
	Symlink(oldname, newname string) error
	Truncate(name string, size int64) error
}
