package text

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Text string

var max_length = 10

func LoadText() Text {
	return Text(loadText())
}

func loadText() string {
	res, err := http.Get("https://raw.githubusercontent.com/monkeytypegame/monkeytype/master/frontend/static/languages/english.json")

	if err != nil {
		fmt.Printf("an error ocurred:%v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("panik")
	}

	words := struct {
		Words []string `json:"words"`
	}{}

	json.Unmarshal(body, &words)

	text := make([]string, 0, max_length)

	for i := 0; i <= max_length; i++ {
		text = append(text, words.Words[i])

	}

	return strings.Join(text, " ")
}
