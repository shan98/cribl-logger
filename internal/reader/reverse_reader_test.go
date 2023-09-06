package reader

import (
	"io"
	"strings"
	"testing"
)

func TestReverseReader(t *testing.T) {
	s := "hi\nhello how are you\nbye"

	rev := NewReverseReader(strings.NewReader(s), len(s), -1)
	rev.chunkSize = 10

	linesFound := []string{}
	for {
		line, _, err := rev.ReadLine()

		if err != nil {
			if err == io.EOF {
				t.Log("end of file")
				break
			} else {
				t.Fatal(err)
				break
			}
		}

		linesFound = append(linesFound, string(line))
	}

	if len(linesFound) != 3 {
		t.Error("Wrong number of lines found: ", len(linesFound))
	}

	if string(linesFound[0]) != "bye" ||
		string(linesFound[1]) != "hello how are you" ||
		string(linesFound[2]) != "hi" {
		t.Error("Wrong lines found: ", linesFound)
	}

}
