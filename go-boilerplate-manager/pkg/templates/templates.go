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

// FetchPublicTemplates fetches the public template list from the given URL.
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
			localTemplates = append(localTemplates, tpl)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk local templates directory %s: %w", dir, err)
	}
	return localTemplates, nil
}
