package app

import (
	"encoding/json"
	"os"
)

type Link struct {
	Enabled bool
	Title   string
	Url     string
}

type Config struct {
	Port        uint
	Address     string
	Assets      string
	Views       string
	Title       string
	Description string
	Nametag     string
	Avatar      bool
	Website     Link
	Links       []Link
}

func ReadJSON(config any, filename string) {
	raw, err := os.ReadFile(filename)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(raw, &config)

	if err != nil {
		panic(err.Error())
	}
}
