package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func LoadJsonConfiguration() (*Config, error) {
	fileFullPath, err := getFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(fileFullPath)
	decoder := json.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to read the json file: %w", err)
	}
	return &config, nil
}

func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to locate home directory: %w", err)
	}
	fileFullPath := filepath.Join(homeDir, configFileName)
	return fileFullPath, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return c.write()
}

func (c *Config) GetUser() string {
	return c.CurrentUserName
}

func (c *Config) write() error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	fileFullPath, err := getFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(fileFullPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil

}
