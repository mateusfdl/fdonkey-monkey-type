package model

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mateusfdl/fdonkey-monkey-type/internal/text"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/window"
)

type model struct {
	Text []rune

	Typed []rune
}

func initialModel() model {
	text := text.LoadText()

	return model{
		Text:  []rune(text),
		Typed: []rune{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		if msg.Type == tea.KeyBackspace {
			m.Typed = m.Typed[:len(m.Typed)-1]
		}

		if len(m.Text) == len(m.Typed) {
			return m, tea.Quit
		}

		next := rune(m.Text[len(m.Typed)])

		if next == '\n' {
			m.Typed = append(m.Typed, next)
		}
		m.Typed = append(m.Typed, msg.Runes...)
	}

	return m, nil
}

func (m model) View() string {
	remaining := string(m.Text[len(m.Typed):])

	var typed string
	for i, c := range m.Typed {
		if c == rune(m.Text[i]) {
			typed += fmt.Sprintf("\x1B[28;2;249;38;114m%v\x1B[0m", string(c))
		} else {
			typed += fmt.Sprintf("\x1B[38;2;249;38;114m%v\x1B[0m", string(m.Text[i]))
		}

	}

	return window.NewWindow().Render(fmt.Sprintf("%s%s", typed, remaining))
}

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("err %v", err)
		os.Exit(1)
	}
}
