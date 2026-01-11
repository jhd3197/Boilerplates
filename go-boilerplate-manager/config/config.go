package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the structure of the config.json file
type Config struct {
	GithubToken   string   `json:"github_token"`
	PrivateRepos  []string `json:"private_repos"`
	PublicRepoURL string   `json:"public_repo_url"`
}

// Prompt represents a single prompt definition within a template
type Prompt struct {
	Label   string `json:"label"`
	Default string `json:"default"`
	Format  string `json:"format"`
}

// RenameRule represents a renaming rule for files
type RenameRule struct {
	Old string `json:"old"`
	New string `json:"new"`
}

// ReplaceValue represents a key-value pair for replacement
type ReplaceValue map[string]string

// ReplaceRule represents a replacement rule for content
type ReplaceRule struct {
	Glob   string       `json:"glob"`
	Values ReplaceValue `json:"values"`
}

// Template represents the structure of a template.json file
type Template struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Prompts     map[string]Prompt `json:"prompts"`
	Rename      map[string]string `json:"rename"` // Simplified for now, as the example shows string to string mapping
	Replace     []ReplaceRule     `json:"replace"`
}

// LoadConfig loads the configuration from the config.json file
func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".boilerplates")
	configFilePath := filepath.Join(configDir, "config.json")

	// Ensure config directory and cache directory exist
	cacheDir := filepath.Join(configDir, "cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	defaultConfig := &Config{
		GithubToken:   "", // Use empty string for nil
		PrivateRepos:  []string{},
		PublicRepoURL: "https://raw.githubusercontent.com/username/repo/main/public-templates.json",
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Config file does not exist, return default config
		return defaultConfig, nil
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config JSON: %w", err)
	}

	// Apply defaults if fields are missing in the loaded config
	if cfg.GithubToken == "" {
		cfg.GithubToken = defaultConfig.GithubToken
	}
	if cfg.PrivateRepos == nil { // json.Unmarshal will leave slice nil if not present
		cfg.PrivateRepos = defaultConfig.PrivateRepos
	}
	if cfg.PublicRepoURL == "" {
		cfg.PublicRepoURL = defaultConfig.PublicRepoURL
	}

	return &cfg, nil
}
