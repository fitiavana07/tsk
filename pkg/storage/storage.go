// Package storage contains all storage-related features.
package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fitiavana07/tsk/pkg/storage/db"
	"github.com/fitiavana07/tsk/pkg/task"
)

// Storage is the generalization of storage driver.
type Storage interface {
	// AddTask adds a new task using the storage driver.
	AddTask(name string) *task.Task

	// Tasks lists all tasks stored in the storage.
	Tasks() []task.Task
}

// NewFileStorage creates a new file storage backend to store data.
func NewFileStorage(file string) *FileStorage {
	jsonData, _ := ioutil.ReadFile(file)
	newDB := &db.DB{}
	json.Unmarshal(jsonData, newDB)

	return &FileStorage{file, *newDB}
}

// FileStorage is backend for storing data into a file.
type FileStorage struct {

	// file is name of the file in which to store the data.
	file string

	// db if the whole database in the file, when loaded in memory.
	db db.DB
}

// IsFirstUse returns whether this is the first use of the app.
// First use is determined by the existence of the file.
func (fs FileStorage) IsFirstUse() bool {
	if _, err := os.Stat(fs.file); os.IsNotExist(err) {
		return true
	}
	return false
}

// Save saves data into the file.
// TODO handle errors on file permissions
func (fs FileStorage) Save() error {
	if !fileExists(fs.file) {
		fl, _ := os.Create(fs.file)
		fl.Close()
	}

	jsonData, _ := json.Marshal(fs.db)
	ioutil.WriteFile(fs.file, jsonData, 0644)

	return nil
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
