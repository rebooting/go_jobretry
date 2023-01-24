package go_diskqueue

import (
	"errors"
	"log"
	"os"
	"testing"
)

// setup files for test cases
func createTestFile(t *testing.T, fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func removeTestFile(t *testing.T, filename string) {
	t.Logf("removing. .. %s\n", filename)
	e := os.Remove(filename)
	if e != nil {
		t.Logf("trying to remove file in cleanup, error %v", e)
	}
}

func TestInsertIntoQueue(t *testing.T) {
	createTestFile(t, "/tmp/dummy.txt")
	removeTestFile(t, "/tmp/dummy.txt")
}

func TestGetExtension(t *testing.T) {
	type TestCase struct {
		Filename    string
		ExpectedErr error
	}

	tc := []TestCase{
		{
			Filename:    "jello.001.pending",
			ExpectedErr: nil,
		},
		{
			Filename:    "jello.pending",
			ExpectedErr: InvalidQueueFormatError{},
		},
	}

	for _, eachCase := range tc {
		createTestFile(t,eachCase.Filename)
		defer removeTestFile(t,eachCase.Filename)
		err := getExtensions(eachCase.Filename)
		if !errors.Is(err, eachCase.ExpectedErr) {
			t.Errorf("expected %v, got %v", eachCase.ExpectedErr, err)
		}
	}
}
