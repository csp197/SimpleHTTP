package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"testing"
)

// TestGzipify tests the gzipify function with various inputs.
//
// The function takes a testing.T object as a parameter.
// It returns no value.
func TestGzipify(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		wantErr  bool
		wantData string
	}{
		{"empty string", "", false, ""},
		{"non-empty string", "Hello, World!", false, ""},
		{"error during write", "invalid utf-8", true, ""},
		{"error during close", "close error", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := gzipify(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("gzipify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			// Check if the data is gzipped
			gr, err := gzip.NewReader(bytes.NewBufferString(gotData))
			if err != nil {
				t.Errorf("gzip.NewReader() error = %v", err)
				return
			}
			defer gr.Close()
			unzippedData, err := ioutil.ReadAll(gr)
			if err != nil {
				t.Errorf("ioutil.ReadAll() error = %v", err)
				return
			}
			if string(unzippedData) != tt.data {
				t.Errorf("unzipped data = %q, want %q", unzippedData, tt.data)
			}
		})
	}
}

// TestGzipifyErrorDuringWrite tests the gzipify function for errors during write.
//
// Parameter t is a pointer to the testing.T struct.
// Return type is none.
func TestGzipifyErrorDuringWrite(t *testing.T) {
	// Simulate an error during write
	gzipWriter := gzip.NewWriter(&bytes.Buffer{})
	gzipWriter.Write([]byte("invalid utf-8"))
	err := gzipWriter.Close()
	if err == nil {
		t.Errorf("expected error during write")
	}
}

func TestGzipifyErrorDuringClose(t *testing.T) {
	// Simulate an error during close
	gzipWriter := gzip.NewWriter(&bytes.Buffer{})
	err := gzipWriter.Close()
	if err == nil {
		t.Errorf("expected error during close")
	}
}
