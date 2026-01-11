package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
	"github.com/jhd3197/go-boilerplate-manager/pkg/templates"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available boilerplates",
	Long:  `This command lists all available boilerplate templates, including local, public, and private repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		fmt.Printf("Configuration loaded: %+v\n", cfg)

		var allTemplates []config.Template

		// Fetch public templates
		publicTemplates, err := templates.FetchPublicTemplates(cfg.PublicRepoURL)
		if err != nil {
			log.Printf("Warning: Could not fetch public templates: %v", err)
		} else {
			allTemplates = append(allTemplates, publicTemplates...)
			fmt.Println("\nPublic Templates:")
			for _, tpl := range publicTemplates {
				fmt.Printf("  - %s (%s): %s\n", tpl.Name, tpl.ID, tpl.Description)
			}
		}

		// Discover local templates
		localTemplates, err := templates.DiscoverLocalTemplates("templates") // Assuming "templates" directory for local ones
		if err != nil {
			log.Printf("Warning: Could not discover local templates: %v", err)
		} else {
			allTemplates = append(allTemplates, localTemplates...)
			fmt.Println("\nLocal Templates:")
			for _, tpl := range localTemplates {
				fmt.Printf("  - %s (%s): %s\n", tpl.Name, tpl.ID, tpl.Description)
			}
		}

		fmt.Println("\nTotal Templates found:", len(allTemplates))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
