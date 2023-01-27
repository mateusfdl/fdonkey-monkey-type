package main

import (
	"flag"

	"github.com/mateusfdl/fdonkey-monkey-type/internal/config"
	"github.com/mateusfdl/fdonkey-monkey-type/internal/model"
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

	model.Start()
}
