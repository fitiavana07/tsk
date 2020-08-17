package persist

import (
	"bytes"
	"encoding/json"
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
