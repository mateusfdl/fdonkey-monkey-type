package config

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
)

type Config struct {
	Theme Theme `mapstructure:"theme"`
}

type Theme struct {
	Background string `mapstructure:"background"`
	Typed      string `mapstructure:"typed"`
	Failed     string `mapstructure:"failed"`
	Font       string `mapstructure:"font_color"`
}

var (
	FilePath = fmt.Sprintf("%v/.config/fdonkey", homeDir())
	FileType = "toml"
	FileName = "theme"
)

func init() {
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)
	viper.AddConfigPath(FilePath)
}

func LoadConfig() Config {
	c := Config{}

	err := viper.ReadInConfig()
	if err != nil {
		c = Config{
			Theme: DefaultTheme(),
		}
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return c
}

func CreateExampleConfigFile() {
	err := os.MkdirAll(FilePath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("%v/%v.%v", FilePath, FileName, FileType))
	if err != nil {
		log.Fatal(err)
	}

	c := Config{
		Theme: DefaultTheme(),
	}

	toml.NewEncoder(f).Encode(c)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func DefaultTheme() Theme {
	return Theme{
		Background: "#1F1F28",
		Typed:      "#999999",
		Failed:     "#ff3a3a",
		Font:       "#bcbcbc",
	}
}

func homeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path
}
