package go_diskqueue

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	PENDING = iota
	DONE
	QUEUE
)

const (
	STATE_QUEUE   = ".queue"
	STATE_PENDING = ".pending"
	STATE_DONE    = ".done"
)

type Queue struct {
}

type QueueExtensions struct {
	stage string
	count int
}

func New() Queue {
	return Queue{}
}

type NotFileError struct{}

func (m NotFileError) Error() string {
	return "Not a file"
}

type InvalidQueueFormatError struct{}

func (m InvalidQueueFormatError) Error() string {
	return "Not a queue file"
}

func InsertIntoQueue(filename string, queueName string) error {
	return nil
}

func getExtensions(filename string) error {
	cleanPath := filepath.Clean(filename)
	//check if filename is a filename
	fileInfo, err := os.Stat(filename)

	if err != nil {
		log.Printf("error when Stat %s", err.Error())
		return err
	}

	if fileInfo.IsDir() {
		return NotFileError{}
	}
	fileStr := filepath.Base(cleanPath)
	fileParts := strings.Split(fileStr, ".")
	if len(fileParts) < 3 {
		return InvalidQueueFormatError{}
	}
	return nil

}
