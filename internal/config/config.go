package config

import (
	"bytes"
	"fmt"
	"log"
	"os"

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
	FilePath     = fmt.Sprintf("%v/.config/fdonkey", homeDir())
	FileType     = "toml"
	FileName     = "theme"
	DefaultTheme = []byte(`[theme]
background = "#3d4349"
fail = "#000"
typed = "#000"
`)
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
		fmt.Println(err)
		err := viper.ReadConfig(bytes.NewBuffer(DefaultTheme))
		if err != nil {
			log.Fatal(err)
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

	err = os.WriteFile(fmt.Sprintf("%v/%v.%v", FilePath, FileName, FileType), DefaultTheme, 0660)
	if err != nil {
		log.Fatal(err)
	}

}

func homeDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path
}
