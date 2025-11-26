package main

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/XPLassal/simple-snake-on-go/render"
)

type Config struct {
	Columns  int  `json:"columns"`
	HardMode bool `json:"hard_mode"`
	UseEmoji bool `json:"use_emoji"`
}

const configFileName = "config.json"

func (cfg *Config) CreateConfig() {
	var cols int
	var hardInput, emojiInput string

	GetNumberOfColumns(&cols)

	fmt.Print("Hard Mode (increase speed)? (y/n): ")
	fmt.Scan(&hardInput)

	fmt.Print("Use Emojis (set 'n' for SSH/Linux)? (y/n): ")
	fmt.Scan(&emojiInput)

	hardMode := hardInput == "y"
	useEmoji := !(emojiInput == "n")

	err := SaveConfig(cols, hardMode, useEmoji)
	if err != nil {
		fmt.Println("Warning: could not save config:", err)
	}

	cfg = &Config{
		Columns:  cols,
		HardMode: hardMode,
		UseEmoji: useEmoji,
	}
}

func LoadConfig() (*Config, bool) {
	file, err := os.Open(configFileName)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, false
	}

	return &cfg, true
}

func SaveConfig(columns int, hardMode bool, useEmoji bool) error {
	cfg := Config{
		Columns:  columns,
		HardMode: hardMode,
		UseEmoji: useEmoji,
	}

	file, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}
