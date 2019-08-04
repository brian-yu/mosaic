package crawl

import (
	"fmt"
)

type Stream struct {
	src       string
	location  string
	insecamId string
}

func Fields() []string {
	slice := make([]string, 3)
	slice[0] = "src"
	slice[1] = "location"
	slice[2] = "insecamId"
	return slice
}

func (s *Stream) String() string {
	return fmt.Sprintf("{src: \"%s\", location: \"%s\", insecamId: \"%s\"}", s.src, s.location, s.insecamId)
}

func (s *Stream) Slice() []string {
	slice := make([]string, 3)
	slice[0] = s.src
	slice[1] = s.location
	slice[2] = s.insecamId
	return slice
}
