package fs_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/codeclysm/fs"
)

func TestWalk(t *testing.T) {
	tmp, _ := ioutil.TempDir("", "")
	chroot := fs.Chroot{
		FS:   fs.OS{},
		Base: tmp,
	}

	io := fs.Base{FS: chroot}

	defer os.RemoveAll(tmp)

	err := io.WriteFile("/file", []byte("hello"), 0644)
	check(t, err)
	err = io.Mkdir("/folder", 0755)
	check(t, err)
	err = io.WriteFile("/folder/file", []byte("hello"), 0644)
	check(t, err)

	files := []string{}

	err = io.Walk("/", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	check(t, err)

	equal(t, fmt.Sprintf("%s", files), "[/ /file /folder /folder/file]")
}

func equal(t *testing.T, got, expected interface{}) {
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected `%v`, got `%v`", expected, got)
	}
}
