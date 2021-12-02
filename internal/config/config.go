package config

import (
	"fmt"
	"os"

	"github.com/ghodss/yaml"
)

type PCIID struct {
	Class  string `json:"class,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Device string `json:"device,omitempty"`
}

func (pciid PCIID) IsValid() bool {
	return pciid.Vendor != "" && pciid.Device != ""
}

type MatchOn struct {
	PCIID PCIID `json:"pciId,omitempty"`
}

type Feature struct {
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	MatchOn []MatchOn `json:"matchOn"`
}

type Config struct {
	Features []Feature `json:"features"`
}

func FromFile(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading %q: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return Config{}, fmt.Errorf("decoding %q: %w", path, err)
	}

	return cfg, nil
}
