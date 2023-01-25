package window

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type Window struct {
	Height        int
	Width         int
	AlignVertical float64
}

func NewWindow() Window {
	widht, height := getSize()

	return Window{
		Height:        height,
		Width:         widht,
		AlignVertical: 0.5,
	}
}

func (w Window) Render(content string) string {
	return lipgloss.NewStyle().
		Width(w.Width).
		Height(w.Height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Position(w.AlignVertical)).
		Render(content)
}

func getSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Printf("err %v", err)
	}
	return w, h
}
