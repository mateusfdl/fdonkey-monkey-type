package main

import (
	"flag"

	"github.com/mateusfdl/fdonkey-monkey-type/internal/config"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/model"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/text"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/theme"
)

var (
	generateFile = flag.Bool("generate-file", false, "generates a config file example")
)

func main() {
	flag.Parse()
	if *generateFile {
		config.CreateExampleConfigFile()
		return
	}

	t := text.LoadText()
	theme := theme.New()

	m := model.Model{
		Text:  []rune(t),
		Typed: []rune{},
		Theme: theme,
	}

	model.Start(m)
}
