package tools

import (
	"io"
	"math"
)

const maxLen = math.MaxInt32 >> 4

type BigSlice struct {
	data       [][]byte
	chunk, pos int
}

func NewBigSlice() *BigSlice {
	return &BigSlice{data: [][]byte{make([]byte, 0)}}
}

func (s *BigSlice) Read(b []byte) (int, error) {
	if s.pos == len(s.data[s.chunk]) {
		s.chunk++
		s.pos = 0
	}
	if s.chunk == len(s.data) {
		return 0, io.EOF
	}
	n := copy(b, s.data[s.chunk][s.pos:])
	s.pos += n
	return n, nil
}

func (s *BigSlice) Write(b []byte) (int, error) {
	room := maxLen - len(s.data[s.chunk])
	if room >= len(b) {
		s.data[s.chunk] = append(s.data[s.chunk][:s.pos], b...)
		s.pos += len(b)
		return len(b), nil
	}
	s.data[s.chunk] = append(s.data[s.chunk][:s.pos], b[:room]...)
	s.data = append(s.data[:s.chunk], append([]byte(nil), b[room:]...))
	s.chunk++
	s.pos = len(b) - room
	return len(b), nil
}

func (s *BigSlice) Reset() {
	s.pos = 0
	s.chunk = 0
}
