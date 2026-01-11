package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
	"github.com/jhd3197/go-boilerplate-manager/pkg/templates"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new boilerplate project",
	Long:  `This command guides you through initializing a new project from an available boilerplate template.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		// Get all templates (public and local)
		var allTemplates []config.Template

		publicTemplates, err := templates.FetchPublicTemplates(cfg.PublicRepoURL)
		if err != nil {
			log.Printf("Warning: Could not fetch public templates: %v", err)
		} else {
			allTemplates = append(allTemplates, publicTemplates...)
		}

		localTemplates, err := templates.DiscoverLocalTemplates("templates")
		if err != nil {
			log.Printf("Warning: Could not discover local templates: %v", err)
		} else {
			allTemplates = append(allTemplates, localTemplates...)
		}

		// Create options for the select field
		var templateOptions []huh.Option[string]
		for _, tpl := range allTemplates {
			templateOptions = append(templateOptions, huh.Option[string]{
				Key: tpl.Name, Value: tpl.ID, // Removed Description field
			})
		}

		// Variables to store user input
		var projectName string
		var authorName string
		var selectedTemplateID string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What is your project name?").
					Description("e.g. My Awesome Project").
					Placeholder("Enter project name").
					Value(&projectName).
					Key("projectName"),

				huh.NewInput().
					Title("What is the author's name?").
					Description("e.g. Your Name").
					Placeholder("Enter author name").
					Value(&authorName).
					Key("authorName"),

				huh.NewSelect[string]().
					Title("Choose a boilerplate template").
					Description("Select from available templates").
					Options(templateOptions...).
					Value(&selectedTemplateID).
					Key("selectedTemplate"),
			),
		).
			WithTheme(huh.ThemeBase())

		err = form.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nProject Name: %s\n", projectName)
		fmt.Printf("Author Name: %s\n", authorName)
		fmt.Printf("Selected Template ID: %s\n", selectedTemplateID)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
