package constants

import (
	_ "embed"
	"encoding/json"
)

//go:embed languages.json
var languagesJson string

func init() {
	err := json.Unmarshal([]byte(languagesJson), &Languages)
	if err != nil {
		panic(err)
	}
}

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Languages []Language
