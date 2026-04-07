package model

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mateusfdl/fdonkey-monkey-type/internal/theme"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/utils"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/window"
)

type Model struct {
	Text []rune

	Typed []rune

	Theme *theme.Theme
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) View() string {
	if len(m.Typed) > len(m.Text) {
		return ""
	}

	var b strings.Builder
	b.Grow(len(m.Text) * 30)

	for i, c := range m.Typed {
		if c == m.Text[i] {
			b.WriteString(utils.Sprintf(m.Theme.Typed, m.Theme.Background, string(c)).String())
		} else {
			b.WriteString(utils.Sprintf(m.Theme.Failed, m.Theme.Background, string(m.Text[i])).String())
		}
	}

	if len(m.Typed) < len(m.Text) {
		remaining := string(m.Text[len(m.Typed):])
		b.WriteString(utils.Sprintf(m.Theme.Font, m.Theme.Background, remaining).String())
	}

	return window.NewWindow().Render(b.String())
}

func Start(m Model) {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("err %v", err)
		os.Exit(1)
	}
}
