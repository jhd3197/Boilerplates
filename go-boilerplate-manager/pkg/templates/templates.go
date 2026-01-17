package templates

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jhd3197/go-boilerplate-manager/config"
)

// FetchRegistry fetches a template registry from the given URL
func FetchRegistry(url string, token string, requiresAuth bool) (*config.TemplateRegistry, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if requiresAuth && token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch registry from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch registry, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var registry config.TemplateRegistry
	if err := json.Unmarshal(body, &registry); err != nil {
		return nil, fmt.Errorf("failed to unmarshal registry JSON: %w", err)
	}

	return &registry, nil
}

// FetchAllRegistries fetches templates from all enabled registries
func FetchAllRegistries(cfg *config.Config) ([]config.Template, error) {
	var allTemplates []config.Template

	for _, reg := range cfg.Registries {
		if !reg.Enabled {
			continue
		}

		registry, err := FetchRegistry(reg.URL, cfg.GithubToken, reg.RequiresAuth)
		if err != nil {
			log.Printf("Warning: Could not fetch registry '%s': %v", reg.ID, err)
			continue
		}

		// Convert registry templates to config templates
		for _, rt := range registry.Templates {
			// Use template's repo if specified, otherwise use registry's default_repo
			repo := registry.DefaultRepo
			if rt.Repo != nil {
				repo = *rt.Repo
			}

			tpl := config.Template{
				ID:          rt.ID,
				Name:        rt.Name,
				Description: rt.Description,
				Category:    rt.Category,
				Repo:        repo,
				Path:        rt.Path,
				Branch:      rt.Branch,
				Commit:      rt.Commit,
				Tags:        rt.Tags,
			}
			allTemplates = append(allTemplates, tpl)
		}
	}

	return allTemplates, nil
}

// GetCustomTemplates converts custom templates from config to Template format
func GetCustomTemplates(cfg *config.Config) []config.Template {
	var templates []config.Template

	for _, ct := range cfg.CustomTemplates {
		tpl := config.Template{
			ID:          ct.ID,
			Name:        ct.Name,
			Description: ct.Description,
			Category:    ct.Category,
			Repo:        ct.Repo,
			Path:        ct.Path,
			Branch:      ct.Branch,
			Commit:      ct.Commit,
			Tags:        ct.Tags,
		}
		templates = append(templates, tpl)
	}

	return templates
}

// FetchPublicTemplates fetches the public template list from the given URL.
// Deprecated: Use FetchRegistry instead
func FetchPublicTemplates(url string) ([]config.Template, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch public templates, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var templates []config.Template
	if err := json.Unmarshal(body, &templates); err != nil {
		return nil, fmt.Errorf("failed to unmarshal public templates JSON: %w", err)
	}

	return templates, nil
}

// DiscoverLocalTemplates walks through the given directory and finds all template.json files.
func DiscoverLocalTemplates(dir string) ([]config.Template, error) {
	var localTemplates []config.Template
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name() == "template.json" {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Warning: Could not read template.json at %s: %v", path, err)
				return nil // Continue walking even if one file fails
			}

			var tpl config.Template
			if err := json.Unmarshal(data, &tpl); err != nil {
				log.Printf("Warning: Could not unmarshal template.json at %s: %v", path, err)
				return nil // Continue walking even if one file fails
			}
			// Store local path for local templates
			tpl.Path = filepath.Dir(path)
			localTemplates = append(localTemplates, tpl)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk local templates directory %s: %w", dir, err)
	}
	return localTemplates, nil
}
