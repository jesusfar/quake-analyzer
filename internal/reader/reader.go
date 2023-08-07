package reader

import (
	"bufio"
	"os"
)

// Reader defines the operation for reading files.
type Reader interface {
	// Read reads a file given a path and return the line chan and error for consume.
	Read(path string) (<-chan string, <-chan error)
}

// FileReader implements Reader interface.
type FileReader struct {
}

// NewFileReader creates a new instance of FileReader.
func NewFileReader() *FileReader {
	return &FileReader{}
}

// Read reads a file given a path and return the line chan and error for consume.
func (f *FileReader) Read(path string) (<-chan string, <-chan error) {
	lineReadChannel := make(chan string)
	errChannel := make(chan error, 1)

	go fileReader(path, lineReadChannel, errChannel)

	return lineReadChannel, errChannel
}

func fileReader(filePath string, lineChannel chan string, errChannel chan error) {
	defer close(lineChannel)
	defer close(errChannel)

	file, err := os.Open(filePath)
	if err != nil {
		errChannel <- err
		_ = file.Close()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineChannel <- line
	}

	if err := scanner.Err(); err != nil {
		errChannel <- err
	}
}
