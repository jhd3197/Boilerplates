package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage boilerplate manager configuration",
	Long:  `View and modify configuration settings including GitHub token and other options.`,
}

// configShowCmd shows the current configuration
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		configDir, _ := config.GetConfigDir()
		fmt.Printf("Config directory: %s\n\n", configDir)

		// GitHub token (masked)
		token := cfg.GithubToken
		if token != "" {
			if len(token) > 8 {
				token = token[:4] + "..." + token[len(token)-4:]
			} else {
				token = "***"
			}
		} else {
			token = "(not set)"
		}
		fmt.Printf("GitHub Token: %s\n\n", token)

		// Registries
		fmt.Printf("Registries (%d):\n", len(cfg.Registries))
		for _, reg := range cfg.Registries {
			status := "enabled"
			if !reg.Enabled {
				status = "disabled"
			}
			fmt.Printf("  - %s (%s): %s\n", reg.ID, status, reg.URL)
		}

		// Custom templates
		fmt.Printf("\nCustom Templates (%d):\n", len(cfg.CustomTemplates))
		for _, tpl := range cfg.CustomTemplates {
			commitInfo := "latest"
			if tpl.Commit != nil {
				commitInfo = (*tpl.Commit)[:min(7, len(*tpl.Commit))]
			}
			fmt.Printf("  - %s [%s]: %s\n", tpl.ID, commitInfo, tpl.Repo)
		}
	},
}

// configSetTokenCmd sets the GitHub token
var configSetTokenCmd = &cobra.Command{
	Use:   "set-token [token]",
	Short: "Set the GitHub personal access token",
	Long: `Set the GitHub personal access token for accessing private repositories.

If no token is provided, you will be prompted to enter it interactively.

Examples:
  boilerplate config set-token ghp_xxxxxxxxxxxx
  boilerplate config set-token  # Interactive prompt`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		var token string
		if len(args) > 0 {
			token = args[0]
		} else {
			// Interactive prompt
			var inputToken string
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("GitHub Personal Access Token").
						Description("Enter your GitHub PAT (starts with ghp_)").
						Placeholder("ghp_xxxxxxxxxxxx").
						Value(&inputToken).
						EchoMode(huh.EchoModePassword),
				),
			)

			if err := form.Run(); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			token = strings.TrimSpace(inputToken)
		}

		cfg.GithubToken = token
		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		if token == "" {
			fmt.Println("GitHub token cleared")
		} else {
			fmt.Println("GitHub token saved successfully")
		}
	},
}

// configClearTokenCmd clears the GitHub token
var configClearTokenCmd = &cobra.Command{
	Use:   "clear-token",
	Short: "Clear the GitHub personal access token",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		cfg.GithubToken = ""
		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Println("GitHub token cleared")
	},
}

// configPathCmd shows the config file path
var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show configuration file path",
	Run: func(cmd *cobra.Command, args []string) {
		configDir, err := config.GetConfigDir()
		if err != nil {
			log.Fatalf("Error getting config directory: %v", err)
		}
		fmt.Printf("Config directory: %s\n", configDir)
		fmt.Printf("Config file: %s/config.json\n", configDir)

		cacheDir, _ := config.GetCacheDir()
		fmt.Printf("Cache directory: %s\n", cacheDir)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetTokenCmd)
	configCmd.AddCommand(configClearTokenCmd)
	configCmd.AddCommand(configPathCmd)
}
