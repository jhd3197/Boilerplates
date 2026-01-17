package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage custom templates",
	Long:  `Add, remove, and manage custom template repositories with optional commit pinning.`,
}

// templateListCmd lists all custom templates
var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all custom templates",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		fmt.Println("Custom Templates:")
		if len(cfg.CustomTemplates) == 0 {
			fmt.Println("  No custom templates configured")
			return
		}

		for _, tpl := range cfg.CustomTemplates {
			commitInfo := "latest"
			if tpl.Commit != nil {
				commitInfo = *tpl.Commit
			}
			private := ""
			if tpl.Private {
				private = " [private]"
			}
			fmt.Printf("  - %s (%s)%s\n", tpl.Name, tpl.ID, private)
			fmt.Printf("    Repo: %s\n", tpl.Repo)
			if tpl.Path != "" {
				fmt.Printf("    Path: %s\n", tpl.Path)
			}
			fmt.Printf("    Commit: %s\n", commitInfo)
			if tpl.Description != "" {
				fmt.Printf("    Description: %s\n", tpl.Description)
			}
		}
	},
}

// templateAddCmd adds a new custom template
var templateAddCmd = &cobra.Command{
	Use:   "add <id> <repo-url>",
	Short: "Add a custom template repository",
	Long: `Add a custom template repository with optional commit pinning.

Examples:
  # Add a template from the repo root
  boilerplate template add my-template https://github.com/user/my-template

  # Add a template from a subdirectory
  boilerplate template add my-template https://github.com/user/templates --path templates/my-template

  # Add a template pinned to a specific commit
  boilerplate template add my-template https://github.com/user/templates --commit abc123

  # Add a private template
  boilerplate template add my-template https://github.com/user/private-templates --private`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		repo := args[1]

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		path, _ := cmd.Flags().GetString("path")
		branch, _ := cmd.Flags().GetString("branch")
		commit, _ := cmd.Flags().GetString("commit")
		private, _ := cmd.Flags().GetBool("private")
		category, _ := cmd.Flags().GetString("category")

		if name == "" {
			name = id
		}
		if branch == "" {
			branch = "main"
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		template := config.CustomTemplate{
			ID:          id,
			Name:        name,
			Description: description,
			Repo:        repo,
			Path:        path,
			Branch:      branch,
			Private:     private,
			Category:    category,
		}

		if commit != "" {
			template.Commit = &commit
		}

		if err := cfg.AddCustomTemplate(template); err != nil {
			log.Fatalf("Error adding template: %v", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Template '%s' added successfully\n", id)
		if template.Commit != nil {
			fmt.Printf("Pinned to commit: %s\n", *template.Commit)
		}
	},
}

// templateRemoveCmd removes a custom template
var templateRemoveCmd = &cobra.Command{
	Use:   "remove <id>",
	Short: "Remove a custom template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		if err := cfg.RemoveCustomTemplate(id); err != nil {
			log.Fatalf("Error removing template: %v", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Template '%s' removed successfully\n", id)
	},
}

// templatePinCmd pins a template to a specific commit
var templatePinCmd = &cobra.Command{
	Use:   "pin <id> <commit>",
	Short: "Pin a custom template to a specific commit",
	Long: `Pin a custom template to a specific commit hash.
This ensures the template always uses that exact version.

Example:
  boilerplate template pin my-template abc123def456`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		commit := args[1]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		if err := cfg.SetTemplateCommit(id, &commit); err != nil {
			log.Fatalf("Error pinning template: %v", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Template '%s' pinned to commit %s\n", id, commit)
	},
}

// templateUnpinCmd removes commit pin from a template
var templateUnpinCmd = &cobra.Command{
	Use:   "unpin <id>",
	Short: "Unpin a custom template (use latest)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		if err := cfg.SetTemplateCommit(id, nil); err != nil {
			log.Fatalf("Error unpinning template: %v", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Template '%s' unpinned, will use latest\n", id)
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.AddCommand(templateListCmd)
	templateCmd.AddCommand(templateAddCmd)
	templateCmd.AddCommand(templateRemoveCmd)
	templateCmd.AddCommand(templatePinCmd)
	templateCmd.AddCommand(templateUnpinCmd)

	templateAddCmd.Flags().String("name", "", "Human-readable name for the template")
	templateAddCmd.Flags().String("description", "", "Description of the template")
	templateAddCmd.Flags().String("path", "", "Path within the repo to the template")
	templateAddCmd.Flags().String("branch", "main", "Branch to use")
	templateAddCmd.Flags().String("commit", "", "Pin to a specific commit hash")
	templateAddCmd.Flags().Bool("private", false, "Template is in a private repository")
	templateAddCmd.Flags().String("category", "", "Category for the template")
}
