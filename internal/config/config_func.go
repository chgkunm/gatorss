package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// helper functions
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error retrieving home directory: %w", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error Opening file[%s]: %w", filePath, err)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(cfg)
	if err != nil {
		return fmt.Errorf("error Encoding to file[%s]: %w", filePath, err)
	}

	return nil
}

// exportable functions
func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("unable to open file: %w", err)
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	jsonBody := Config{}
	if err := decoder.Decode(&jsonBody); err != nil {
		return Config{}, fmt.Errorf("error Decoding from file[%s]: %w", filePath, err)
	}

	return jsonBody, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.Current_user_name = userName
	return write(*cfg)
}
