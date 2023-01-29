package theme

import "github.com/mateusfdl/fdonkey-monkey-type/internal/config"

type Theme struct {
	Background string
	Typed      string
	Failed     string
	Font       string
}

func New() *Theme {
	c := config.LoadConfig()
	return &Theme{
		Background: c.Theme.Background,
		Typed:      c.Theme.Typed,
		Failed:     c.Theme.Failed,
		Font:       c.Theme.Font,
	}
}
