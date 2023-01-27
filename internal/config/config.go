package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config interface {
	Get(field string) interface{}
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
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper.GetViper()
}

func CreateExampleConfigFile() {
	err := os.MkdirAll(FilePath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	err = os.WriteFile(fmt.Sprintf("%v/%v.%v", FilePath, FileName, FileType), exampleConfigFile(), 0660)
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

func exampleConfigFile() []byte {
	return []byte(`[theme]
background = "#fff"
success = "#000"
fail = "#000"
typed = "#000"
`)
}
