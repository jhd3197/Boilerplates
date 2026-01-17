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

		// Get all templates
		var allTemplates []config.Template

		// Fetch from registries
		registryTemplates, err := templates.FetchAllRegistries(cfg)
		if err != nil {
			log.Printf("Warning: Could not fetch registry templates: %v", err)
		} else {
			allTemplates = append(allTemplates, registryTemplates...)
		}

		// Get custom templates
		customTemplates := templates.GetCustomTemplates(cfg)
		allTemplates = append(allTemplates, customTemplates...)

		// Discover local templates
		localTemplates, err := templates.DiscoverLocalTemplates("templates")
		if err != nil {
			log.Printf("Warning: Could not discover local templates: %v", err)
		} else {
			allTemplates = append(allTemplates, localTemplates...)
		}

		if len(allTemplates) == 0 {
			fmt.Println("No templates available. Add a registry or custom template first.")
			return
		}

		// Create options for the select field
		var templateOptions []huh.Option[string]
		for _, tpl := range allTemplates {
			label := fmt.Sprintf("%s (%s)", tpl.Name, tpl.Category)
			templateOptions = append(templateOptions, huh.Option[string]{
				Key: label, Value: tpl.ID,
			})
		}

		// Variables to store user input
		var projectName string
		var authorName string
		var selectedTemplateID string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Choose a boilerplate template").
					Description("Select from available templates").
					Options(templateOptions...).
					Value(&selectedTemplateID).
					Key("selectedTemplate"),

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
			),
		).
			WithTheme(huh.ThemeBase())

		err = form.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Find selected template
		var selectedTemplate *config.Template
		for _, tpl := range allTemplates {
			if tpl.ID == selectedTemplateID {
				selectedTemplate = &tpl
				break
			}
		}

		if selectedTemplate == nil {
			fmt.Printf("Error: Template '%s' not found\n", selectedTemplateID)
			os.Exit(1)
		}

		fmt.Printf("\nProject Name: %s\n", projectName)
		fmt.Printf("Author Name: %s\n", authorName)
		fmt.Printf("Selected Template: %s (%s)\n", selectedTemplate.Name, selectedTemplate.ID)
		if selectedTemplate.Repo != "" {
			fmt.Printf("Repo: %s\n", selectedTemplate.Repo)
			fmt.Printf("Path: %s\n", selectedTemplate.Path)
		}

		// TODO: Implement actual template cloning and variable replacement
		fmt.Println("\nTemplate cloning not yet implemented. Use 'create' command for now.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
