// Package storage contains all storage-related features.
package storage

import (
	"os"
)

// TODO load data from rw, write "FileVersion": "1" to mark a first visit

// NewFileStorage creates a new file storage backend to store data
func NewFileStorage(file string) *FileStorage {
	return &FileStorage{file}
}

// FileStorage is backend for storing data into a file
type FileStorage struct {
	file string
}

// IsFirstUse returns whether this is the first use of the app.
// First use is determined by the existence of the file.
func (fs FileStorage) IsFirstUse() bool {
	if _, err := os.Stat(fs.file); os.IsNotExist(err) {
		return true
	}
	return false
}

// Save saves data into the file,
func (fs FileStorage) Save() error {
	fl, err := os.Create(fs.file)
	if err != nil {
		return err
	}
	fl.Close()
	return nil
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
