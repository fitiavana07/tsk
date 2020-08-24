package storage

import (
	"testing"

	"os"
	"path/filepath"
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
