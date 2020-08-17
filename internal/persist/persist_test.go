package persist

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteData(t *testing.T) {
	// buffer to get the data to write
	buffer := &bytes.Buffer{}

	// given: data,
	data := struct{ Hey string }{"value"}

	// when: I write to a file
	WriteData(data, buffer)

	// then: file contains the content in JSON format
	got := buffer.Bytes()
	want, _ := json.Marshal(data)

	if bytes.Compare(got, want) != 0 {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestReadData(t *testing.T) {
	// given: the data
	data := &struct{ Key string }{"value"}

	// given: buffer containing the JSON data to read from
	content, _ := json.Marshal(data)
	dataBuffer := bytes.NewBuffer(content)

	// when: I read from the file
	got := &struct{ Key string }{}
	ReadData(dataBuffer, got)

	// then: the read content corresponds to the data

	if *data != *got {
		t.Errorf("got %+v, want %+v", got, data)
	}
}

func TestWriteDataToFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "FILE.JSON")

	// given: data
	data := struct{ Key string }{"value"}

	// when: I write to a file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Errorf("%v", err)
	}
	WriteData(data, file)

	// then: the file is not empty
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(content) == 0 {
		t.Errorf("created file %q is empty", filePath)
	}
}

func TestReadDataFromFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "FILE.JSON")

	writtenFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	data := struct{ Key string }{"value"}
	WriteData(data, writtenFile)

	got := &struct{ Key string }{}
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("%v", err)
	}
	ReadData(file, got)
	if *got != data {
		t.Errorf("got %v, wanted %v", got, data)
	}
}
