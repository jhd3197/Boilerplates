package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
	"github.com/jhd3197/go-boilerplate-manager/pkg/templates"
)

var (
	projectName string
	authorName  string
	templateID  string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project from a boilerplate template",
	Long:  `This command creates a new project from a boilerplate template with the provided arguments.`, 
	Run: func(cmd *cobra.Command, args []string) {
		// Find the selected template
		localTemplates, err := templates.DiscoverLocalTemplates("templates")
		if err != nil {
			log.Fatalf("Error discovering local templates: %v", err)
		}

		var selectedTemplate *config.Template
		for i := range localTemplates {
			if localTemplates[i].ID == templateID {
				selectedTemplate = &localTemplates[i]
				break
			}
		}

		if selectedTemplate == nil {
			log.Fatalf("Template with ID '%s' not found", templateID)
		}

		fmt.Printf("Creating project '%s' from template '%s'\n", projectName, selectedTemplate.Name)

		// Create project directory
		destDir := projectName
		if err := os.MkdirAll(destDir, 0755); err != nil {
			log.Fatalf("Failed to create project directory: %v", err)
		}

		// Template data
		templateData := map[string]string{
			"ProjectName": projectName,
			"AuthorName":  authorName,
		}

		// Walk through the template directory
		templateDir := filepath.Join("templates", selectedTemplate.ID)
		err = filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip the template.json file
			if info.Name() == "template.json" {
				return nil
			}

			relPath, err := filepath.Rel(templateDir, path)
			if err != nil {
				return err
			}

			destPath := filepath.Join(destDir, relPath)

			// Rename file if there is a rule for it
			for oldName, newNameTpl := range selectedTemplate.Rename {
				if info.Name() == oldName {
					var newName bytes.Buffer
					nameTpl, err := template.New("newName").Parse(newNameTpl)
					if err != nil {
						return fmt.Errorf("failed to parse rename template: %w", err)
					}
					if err := nameTpl.Execute(&newName, templateData); err != nil {
						return fmt.Errorf("failed to execute rename template: %w", err)
					}
					destPath = filepath.Join(filepath.Dir(destPath), newName.String())
					break
				}
			}

			if info.IsDir() {
				return os.MkdirAll(destPath, info.Mode())
			}

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Apply template to file content
			tmpl, err := template.New("file").Parse(string(content))
			if err != nil {
				return fmt.Errorf("failed to parse file template %s: %w", path, err)
			}

			var processedContent bytes.Buffer
			if err := tmpl.Execute(&processedContent, templateData); err != nil {
				return fmt.Errorf("failed to execute file template %s: %w", path, err)
			}

			return os.WriteFile(destPath, processedContent.Bytes(), info.Mode())
		})

		if err != nil {
			log.Fatalf("Failed to create project: %v", err)
		}

		fmt.Println("Project created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&projectName, "project-name", "", "The name of the project")
	createCmd.Flags().StringVar(&authorName, "author-name", "", "The name of the author")
	createCmd.Flags().StringVar(&templateID, "template-id", "", "The ID of the template to use")
	createCmd.MarkFlagRequired("project-name")
	createCmd.MarkFlagRequired("author-name")
	createCmd.MarkFlagRequired("template-id")
}