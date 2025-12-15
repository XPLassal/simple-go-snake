package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/XPLassal/simple-go-snake/render"
)

type Config struct {
	Columns        int  `json:"columns"`
	HardMode       bool `json:"hard_mode"`
	UseEmoji       bool `json:"use_emoji"`
	FPS            int  `json:"fps"`
	AllowWallPass  bool `json:"allow_wall_pass"`
	MultiAppleMode bool `json:"multi_apple_mode"`
}

const appDirName = "simple-go-snake"
const configFileName = "config.json"

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, appDirName, configFileName), nil
}

func askQuation[T any](s string, answer *T) {
	fmt.Print(s)
	fmt.Scan(answer)
}

func CreateConfig() *Config {
	var cols, fps int
	var answer string

	GetNumberOfColumns(&cols)

	for fps <= 0 {
		askQuation("Set FPS number (10 recommended): ", &fps)
	}

	askQuation("Allow the snake to pass through walls? (y/n): ", &answer)
	allowWallPass := answer == "y"

	askQuation("Allow multiple apples to spawn? (y/n): ", &answer)
	multiAppleMode := answer == "y"

	askQuation("Hard Mode (increase speed)? (y/n): ", &answer)
	hardMode := answer == "y"

	askQuation("Use Emojis (set 'n' for SSH/Linux)? (y/n): ", &answer)
	useEmoji := !(answer == "n")

	err := SaveConfig(cols, hardMode, useEmoji, fps, allowWallPass, multiAppleMode)
	if err != nil {
		path, _ := getConfigPath()
		fmt.Printf("Warning: could not save config to %s: %v\n", path, err)
	}

	return &Config{
		Columns:        cols,
		HardMode:       hardMode,
		UseEmoji:       useEmoji,
		FPS:            fps,
		AllowWallPass:  allowWallPass,
		MultiAppleMode: multiAppleMode,
	}
}

func LoadConfig() (*Config, bool) {
	path, err := getConfigPath()
	if err != nil {
		return nil, false
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, false
	}

	if err := ValidateConfig(&cfg); err != nil {
		fmt.Printf("Error: %v. Let's regenerate.\n", err)
		cfg = *CreateConfig()
	}

	return &cfg, true
}

func ValidateConfig(c *Config) error {
	if c.Columns == 0 || c.FPS == 0 {
		return errors.New("incorrect config")
	}
	return nil
}

func SaveConfig(columns int, hardMode bool, useEmoji bool, fps int, allowWallPass bool, multiAppleMode bool) error {
	cfg := Config{
		Columns:        columns,
		HardMode:       hardMode,
		UseEmoji:       useEmoji,
		FPS:            fps,
		AllowWallPass:  allowWallPass,
		MultiAppleMode: multiAppleMode,
	}

	path, err := getConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}
