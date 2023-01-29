package utils

import (
	"math/rand"
	"time"

	"github.com/muesli/termenv"
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

func Sprintf(fg, bg, s string) termenv.Style {
	o := termenv.String(s)

	o = o.Foreground(termenv.RGBColor(fg))
	o = o.Background(termenv.RGBColor(bg))

	return o
}
