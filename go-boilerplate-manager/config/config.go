package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Registry represents a template registry source
type Registry struct {
	ID           string `json:"id"`
	Name         string `json:"name,omitempty"`
	URL          string `json:"url"`
	Enabled      bool   `json:"enabled"`
	RequiresAuth bool   `json:"requires_auth,omitempty"`
}

// CustomTemplate represents a user-added template with optional commit pinning
type CustomTemplate struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Repo        string  `json:"repo"`
	Path        string  `json:"path,omitempty"`
	Branch      string  `json:"branch,omitempty"`
	Commit      *string `json:"commit,omitempty"` // nil means latest, string means pinned
	Private     bool    `json:"private,omitempty"`
	Category    string  `json:"category,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// Config represents the structure of the config.json file
type Config struct {
	GithubToken     string           `json:"github_token"`
	Registries      []Registry       `json:"registries"`
	CustomTemplates []CustomTemplate `json:"custom_templates"`
	// Legacy fields for backward compatibility
	PrivateRepos  []string `json:"private_repos,omitempty"`
	PublicRepoURL string   `json:"public_repo_url,omitempty"`
}

// Prompt represents a single prompt definition within a template
type Prompt struct {
	Label   string `json:"label"`
	Default string `json:"default"`
	Format  string `json:"format"`
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
	Category    string            `json:"category,omitempty"`
	Prompts     map[string]Prompt `json:"prompts"`
	Rename      map[string]string `json:"rename"`
	Replace     []ReplaceRule     `json:"replace"`
	// Registry fields
	Repo   string   `json:"repo,omitempty"`
	Path   string   `json:"path,omitempty"`
	Branch string   `json:"branch,omitempty"`
	Commit *string  `json:"commit,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

// TemplateRegistry represents the structure of a templates-registry.json file
type TemplateRegistry struct {
	Version     string             `json:"version"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	DefaultRepo string             `json:"default_repo"` // Used when template.repo is null
	Templates   []RegistryTemplate `json:"templates"`
}

// RegistryTemplate represents a template entry in the registry
type RegistryTemplate struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Repo        *string  `json:"repo"` // nil means use registry's default_repo
	Path        string   `json:"path"`
	Branch      string   `json:"branch"`
	Commit      *string  `json:"commit"`
	Tags        []string `json:"tags"`
}

// GetConfigDir returns the path to the config directory
func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".boilerplates"), nil
}

// GetCacheDir returns the path to the cache directory
func GetCacheDir() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "cache"), nil
}

// LoadConfig loads the configuration from the config.json file
func LoadConfig() (*Config, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	configFilePath := filepath.Join(configDir, "config.json")

	// Ensure config directory and cache directory exist
	cacheDir := filepath.Join(configDir, "cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	defaultConfig := &Config{
		GithubToken: "",
		Registries: []Registry{
			{
				ID:      "official",
				Name:    "Official Templates",
				URL:     "https://raw.githubusercontent.com/jgarzarebel/boilerplates/main/templates-registry.json",
				Enabled: true,
			},
		},
		CustomTemplates: []CustomTemplate{},
		PrivateRepos:    []string{},
		PublicRepoURL:   "",
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
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

	// Apply defaults if fields are missing
	if cfg.Registries == nil {
		cfg.Registries = defaultConfig.Registries
	}
	if cfg.CustomTemplates == nil {
		cfg.CustomTemplates = []CustomTemplate{}
	}

	return &cfg, nil
}

// SaveConfig saves the configuration to the config.json file
func SaveConfig(cfg *Config) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configFilePath := filepath.Join(configDir, "config.json")

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config JSON: %w", err)
	}

	if err := os.WriteFile(configFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// AddRegistry adds a new registry to the config
func (c *Config) AddRegistry(registry Registry) error {
	for _, r := range c.Registries {
		if r.ID == registry.ID {
			return fmt.Errorf("registry with ID '%s' already exists", registry.ID)
		}
	}
	c.Registries = append(c.Registries, registry)
	return nil
}

// RemoveRegistry removes a registry from the config by ID
func (c *Config) RemoveRegistry(id string) error {
	for i, r := range c.Registries {
		if r.ID == id {
			c.Registries = append(c.Registries[:i], c.Registries[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("registry with ID '%s' not found", id)
}

// AddCustomTemplate adds a new custom template to the config
func (c *Config) AddCustomTemplate(template CustomTemplate) error {
	for _, t := range c.CustomTemplates {
		if t.ID == template.ID {
			return fmt.Errorf("template with ID '%s' already exists", template.ID)
		}
	}
	c.CustomTemplates = append(c.CustomTemplates, template)
	return nil
}

// RemoveCustomTemplate removes a custom template from the config by ID
func (c *Config) RemoveCustomTemplate(id string) error {
	for i, t := range c.CustomTemplates {
		if t.ID == id {
			c.CustomTemplates = append(c.CustomTemplates[:i], c.CustomTemplates[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("template with ID '%s' not found", id)
}

// SetTemplateCommit sets or clears the commit pin for a custom template
func (c *Config) SetTemplateCommit(id string, commit *string) error {
	for i, t := range c.CustomTemplates {
		if t.ID == id {
			c.CustomTemplates[i].Commit = commit
			return nil
		}
	}
	return fmt.Errorf("template with ID '%s' not found", id)
}
