package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
	"github.com/jhd3197/go-boilerplate-manager/pkg/templates"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import <registry-url>",
	Short: "Import templates from a registry URL",
	Long: `Import one or more templates from a remote registry URL.

This allows you to browse and selectively import templates from any registry JSON file.

Examples:
  # Import from a registry URL (lists available templates)
  boilerplate import https://raw.githubusercontent.com/user/templates/main/registry.json

  # Import a specific template by ID
  boilerplate import https://raw.githubusercontent.com/user/templates/main/registry.json --template my-template

  # Import all templates from a registry
  boilerplate import https://raw.githubusercontent.com/user/templates/main/registry.json --all`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		registryURL := args[0]
		templateID, _ := cmd.Flags().GetString("template")
		importAll, _ := cmd.Flags().GetBool("all")
		requiresAuth, _ := cmd.Flags().GetBool("auth")

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		// Fetch the registry
		fmt.Printf("Fetching registry from %s...\n", registryURL)
		registry, err := templates.FetchRegistry(registryURL, cfg.GithubToken, requiresAuth)
		if err != nil {
			log.Fatalf("Error fetching registry: %v", err)
		}

		fmt.Printf("Registry: %s (%d templates)\n", registry.Name, len(registry.Templates))

		if !importAll && templateID == "" {
			// Just list available templates
			fmt.Println("\nAvailable templates:")
			for _, tpl := range registry.Templates {
				repo := registry.DefaultRepo
				if tpl.Repo != nil {
					repo = *tpl.Repo
				}
				fmt.Printf("  - %s: %s\n", tpl.ID, tpl.Name)
				fmt.Printf("    Repo: %s\n", repo)
				fmt.Printf("    Path: %s\n", tpl.Path)
			}
			fmt.Println("\nUse --template <id> to import a specific template, or --all to import all.")
			return
		}

		var templatesToImport []config.RegistryTemplate
		if importAll {
			templatesToImport = registry.Templates
		} else {
			for _, tpl := range registry.Templates {
				if tpl.ID == templateID {
					templatesToImport = append(templatesToImport, tpl)
					break
				}
			}
			if len(templatesToImport) == 0 {
				log.Fatalf("Template '%s' not found in registry", templateID)
			}
		}

		// Import templates as custom templates
		imported := 0
		for _, rt := range templatesToImport {
			repo := registry.DefaultRepo
			if rt.Repo != nil {
				repo = *rt.Repo
			}

			customTpl := config.CustomTemplate{
				ID:          rt.ID,
				Name:        rt.Name,
				Description: rt.Description,
				Repo:        repo,
				Path:        rt.Path,
				Branch:      rt.Branch,
				Commit:      rt.Commit,
				Category:    rt.Category,
				Tags:        rt.Tags,
				Private:     requiresAuth,
			}

			if err := cfg.AddCustomTemplate(customTpl); err != nil {
				log.Printf("Skipping '%s': %v", rt.ID, err)
				continue
			}
			imported++
			fmt.Printf("Imported: %s\n", rt.ID)
		}

		if imported > 0 {
			if err := config.SaveConfig(cfg); err != nil {
				log.Fatalf("Error saving configuration: %v", err)
			}
			fmt.Printf("\nSuccessfully imported %d template(s)\n", imported)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().String("template", "", "Import a specific template by ID")
	importCmd.Flags().Bool("all", false, "Import all templates from the registry")
	importCmd.Flags().Bool("auth", false, "Registry requires authentication")
}
