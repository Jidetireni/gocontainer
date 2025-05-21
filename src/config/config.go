package config

import (
    "encoding/json"
    "os"
)

// Config holds the configuration settings for the container.
type Config struct {
    Rootfs      string `json:"rootfs"`
    Cgroup      string `json:"cgroup"`
    Namespace   string `json:"namespace"`
    Network     string `json:"network"`
}

// LoadConfig reads the configuration from a JSON file.
func LoadConfig(filePath string) (*Config, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config Config
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }

    return &config, nil
}