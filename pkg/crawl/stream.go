package crawl

import (
	"fmt"
)

type Stream struct {
	src       string
	location  string
	insecamId string
}

// Fields returns slice representation of field labels.
func Fields() []string {
	slice := make([]string, 3)
	slice[0] = "src"
	slice[1] = "location"
	slice[2] = "insecamId"
	return slice
}

// String returns a string representation of Stream.
func (s *Stream) String() string {
	return fmt.Sprintf("{src: \"%s\", location: \"%s\", insecamId: \"%s\"}", s.src, s.location, s.insecamId)
}

// Slice returns a slice representation of Stream.
func (s *Stream) Slice() []string {
	slice := make([]string, 3)
	slice[0] = s.src
	slice[1] = s.location
	slice[2] = s.insecamId
	return slice
}
