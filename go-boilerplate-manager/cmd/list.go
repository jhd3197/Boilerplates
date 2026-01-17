package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
	"github.com/jhd3197/go-boilerplate-manager/pkg/templates"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available boilerplates",
	Long:  `This command lists all available boilerplate templates from registries, custom templates, and local sources.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		var allTemplates []config.Template

		// Fetch from registries
		fmt.Println("Fetching templates from registries...")
		registryTemplates, err := templates.FetchAllRegistries(cfg)
		if err != nil {
			log.Printf("Warning: Error fetching registries: %v", err)
		} else if len(registryTemplates) > 0 {
			allTemplates = append(allTemplates, registryTemplates...)
			fmt.Printf("\nRegistry Templates (%d):\n", len(registryTemplates))
			printTemplatesByCategory(registryTemplates)
		}

		// Get custom templates
		customTemplates := templates.GetCustomTemplates(cfg)
		if len(customTemplates) > 0 {
			allTemplates = append(allTemplates, customTemplates...)
			fmt.Printf("\nCustom Templates (%d):\n", len(customTemplates))
			for _, tpl := range customTemplates {
				commitInfo := "latest"
				if tpl.Commit != nil {
					commitInfo = (*tpl.Commit)[:7] // Show short commit hash
				}
				fmt.Printf("  - %s (%s) [%s]: %s\n", tpl.Name, tpl.ID, commitInfo, tpl.Description)
			}
		}

		// Discover local templates
		localTemplates, err := templates.DiscoverLocalTemplates("templates")
		if err != nil {
			log.Printf("Warning: Could not discover local templates: %v", err)
		} else if len(localTemplates) > 0 {
			allTemplates = append(allTemplates, localTemplates...)
			fmt.Printf("\nLocal Templates (%d):\n", len(localTemplates))
			for _, tpl := range localTemplates {
				fmt.Printf("  - %s (%s): %s\n", tpl.Name, tpl.ID, tpl.Description)
			}
		}

		fmt.Printf("\nTotal Templates: %d\n", len(allTemplates))
	},
}

func printTemplatesByCategory(templates []config.Template) {
	categories := make(map[string][]config.Template)
	for _, tpl := range templates {
		cat := tpl.Category
		if cat == "" {
			cat = "other"
		}
		categories[cat] = append(categories[cat], tpl)
	}

	for cat, tpls := range categories {
		fmt.Printf("  [%s]\n", strings.ToUpper(cat))
		for _, tpl := range tpls {
			commitInfo := "latest"
			if tpl.Commit != nil {
				commitInfo = (*tpl.Commit)[:7]
			}
			tags := ""
			if len(tpl.Tags) > 0 {
				tags = " #" + strings.Join(tpl.Tags, " #")
			}
			fmt.Printf("    - %s (%s) [%s]: %s%s\n", tpl.Name, tpl.ID, commitInfo, tpl.Description, tags)
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
