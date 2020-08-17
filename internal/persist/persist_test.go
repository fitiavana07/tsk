package persist

import (
	"bytes"
	"encoding/json"
	//"io/ioutil"
	"testing"
)

func TestWriteToFile(t *testing.T) {
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
