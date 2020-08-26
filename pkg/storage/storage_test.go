package storage

import (
	"testing"

	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fitiavana07/tsk/pkg/storage/db"
	"github.com/fitiavana07/tsk/pkg/task"
)

func tempFile(t *testing.T) string {
	return filepath.Join(t.TempDir(), "data.db")
}

func TestIsFirstUse(t *testing.T) {
	file := tempFile(t)
	fs := NewFileStorage(file)

	t.Run("FirstUse", func(t *testing.T) {
		if !fs.IsFirstUse() {
			t.Errorf("got fs.IsFirstUse() == false, want true")
		}
	})

	t.Run("SecondUse", func(t *testing.T) {
		fl, err := os.Create(file)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		defer fl.Close()

		if fs.IsFirstUse() {
			t.Errorf("got fs.IsFirstUse() == true, want false")
		}
	})
}

func TestSave(t *testing.T) {
	file := tempFile(t)
	fs := NewFileStorage(file)
	fs.Save()

	if !fileExists(file) {
		t.Errorf("file still doesn't exist")
	}
}

func TestSaveNotEmpty(t *testing.T) {
	file := tempFile(t)
	fs := FileStorage{file, db.DB{
		LastID: 1,
		Tasks: []task.Task{
			task.Task{
				ID:   1,
				Name: "a sample task",
			},
		},
	}}
	fs.Save()
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(fileContent) == 0 {
		t.Errorf("file is still empty")
	}
}

func TestNewFileStorageNotEmpty(t *testing.T) {
	file := tempFile(t)
	fs := FileStorage{file, db.DB{
		LastID: 1,
		Tasks: []task.Task{
			task.Task{
				ID:   1,
				Name: "a sample task",
			},
		},
	}}
	fs.Save()

	newFs := NewFileStorage(file)
	got := newFs.db.LastID
	if got != 1 {
		t.Errorf("db.LastID=%d, want=1", got)
	}
}

func TestFileExists(t *testing.T) {
	file := tempFile(t)
	t.Run("FileDoesntExist", func(t *testing.T) {
		if fileExists(file) {
			t.Errorf("got fileExists() == true, want false")
		}
	})
	t.Run("FileExists", func(t *testing.T) {
		fl, err := os.Create(file)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		defer fl.Close()

		if !fileExists(file) {
			t.Errorf("got fileExists() == false, want true")
		}
	})
}
