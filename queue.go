package go_diskqueue

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	PENDING = iota
	DONE
	QUEUE
)

const RETRY_INFO = 1
const STATE_INFO = 2
const FILE_INFO = 0

const (
	STATE_QUEUE   = "queue"
	STATE_PENDING = "pending"
	STATE_DONE    = "done"
)

type QueueItem struct {
	Name       string
	RetryCount int
	State      string
}

// creates new QueueItem representing the actual file
func NewQueueItem(name string, retryCount int, state string) QueueItem {
	return QueueItem{Name: name, RetryCount: retryCount, State: state}
}

type QueueExtensions struct {
	stage string
	count int
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

func ExceedRetries(qi QueueItem, maxRetries int) bool {
	return qi.RetryCount >= maxRetries
}
func getExtensions(filename string) (QueueItem, error) {
	cleanPath := filepath.Clean(filename)
	//check if filename is a filename
	fileInfo, err := os.Stat(filename)

	if err != nil {
		log.Printf("error when Stat %s", err.Error())
		return QueueItem{}, err
	}

	if fileInfo.IsDir() {
		return QueueItem{}, NotFileError{}
	}
	fileStr := filepath.Base(cleanPath)
	fileParts := strings.Split(fileStr, ".")
	if len(fileParts) < 3 {
		return QueueItem{}, InvalidQueueFormatError{}
	}
	retryInfo, err := strconv.Atoi(fileParts[RETRY_INFO])
	if err != nil {
		log.Printf("Queue info of fileParts %v not valid, error: %s", fileParts, err)
		return QueueItem{}, InvalidQueueFormatError{}
	}

	if fileParts[STATE_INFO] != STATE_DONE &&
		fileParts[STATE_INFO] != STATE_PENDING &&
		fileParts[STATE_INFO] != STATE_QUEUE {
		log.Printf("Invalid state for %s, found %s", filename, fileParts[STATE_INFO])
		return QueueItem{}, InvalidQueueFormatError{}
	}
	return NewQueueItem(fileParts[FILE_INFO], retryInfo, fileParts[STATE_INFO]), nil

}
