package reader

import (
	"bytes"
	"io"
)

type ReverseReader struct {
	reader      io.ReaderAt
	data        []byte
	chunkedData []byte
	position    int
	prePosition int // used to track when we've reached the last line (ie. the top most line)
	linesToRead int
	linesRead   int
	chunkSize   int
}

var DefaultChunkSize = 1000

func NewReverseReader(r io.ReaderAt, from, linesToRead int) *ReverseReader {
	rev := &ReverseReader{
		reader:      r,
		position:    from,
		linesToRead: linesToRead,
		chunkSize:   DefaultChunkSize,
	}

	return rev
}

func (r *ReverseReader) ReadLine() ([]byte, int, error) {
	for {
		start := bytes.LastIndexByte(r.data, '\n')
		if start == -1 {
			err := r.loadMoreData()
			if err != nil {
				if err == io.EOF && len(r.data) > 0 {
					line := r.data
					r.data = []byte{} // empty out data since we've now read the last line
					return line, 0, nil
				}

				return nil, 0, err
			}
			continue
		}

		line := r.data[start+1:]
		r.data = r.data[:start]

		return line, r.position, nil
	}
}

func (r *ReverseReader) loadMoreData() error {
	if r.position == 0 {
		return io.EOF
	}

	r.prePosition = r.position
	// the position in the reader starts off from the end
	r.position -= r.chunkSize

	if r.position < 0 {
		r.position = 0
	}

	size := r.chunkSize
	if (r.prePosition - r.position) < r.chunkSize {
		size = r.prePosition - r.position
	}

	r.chunkedData = make([]byte, size)

	_, err := r.reader.ReadAt(r.chunkedData, int64(r.position))

	if err != nil && err != io.EOF {
		return err
	}
	r.data = append(r.chunkedData, r.data...)

	return nil
}
