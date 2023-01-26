package utils

import (
	"math/rand"
	"time"

	"github.com/thefuga/go-collections"
)

type Text struct {
	Words []string
}

func New(s []string) *Text {
	return &Text{
		Words: s,
	}
}

func (s *Text) Shuffle() *Text {
	rand.Seed(time.Now().UnixMicro())

	s.Words = collections.Shuffle(s.Words)

	return s
}

func (s *Text) Chunck(size int) *Text {
	s.Words = collections.ForPage(s.Words, 1, size)

	return s
}
