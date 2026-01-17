package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/jhd3197/go-boilerplate-manager/config"
)

// registryCmd represents the registry command
var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage template registries",
	Long:  `Add, remove, list, and manage template registries and custom templates.`,
}

// registryListCmd lists all registries
var registryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured registries",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		fmt.Println("Configured Registries:")
		if len(cfg.Registries) == 0 {
			fmt.Println("  No registries configured")
			return
		}

		for _, reg := range cfg.Registries {
			status := "enabled"
			if !reg.Enabled {
				status = "disabled"
			}
			auth := ""
			if reg.RequiresAuth {
				auth = " [requires auth]"
			}
			fmt.Printf("  - %s (%s)%s\n    URL: %s\n", reg.ID, status, auth, reg.URL)
		}
	},
}

// registryAddCmd adds a new registry
var registryAddCmd = &cobra.Command{
	Use:   "add <id> <url>",
	Short: "Add a new template registry",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		url := args[1]

		requiresAuth, _ := cmd.Flags().GetBool("auth")
		name, _ := cmd.Flags().GetString("name")

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		registry := config.Registry{
			ID:           id,
			Name:         name,
			URL:          url,
			Enabled:      true,
			RequiresAuth: requiresAuth,
		}

		if err := cfg.AddRegistry(registry); err != nil {
			log.Fatalf("Error adding registry: %v", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Registry '%s' added successfully\n", id)
	},
}

// registryRemoveCmd removes a registry
var registryRemoveCmd = &cobra.Command{
	Use:   "remove <id>",
	Short: "Remove a template registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		if err := cfg.RemoveRegistry(id); err != nil {
			log.Fatalf("Error removing registry: %v", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Registry '%s' removed successfully\n", id)
	},
}

// registryEnableCmd enables a registry
var registryEnableCmd = &cobra.Command{
	Use:   "enable <id>",
	Short: "Enable a template registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		found := false
		for i, reg := range cfg.Registries {
			if reg.ID == id {
				cfg.Registries[i].Enabled = true
				found = true
				break
			}
		}

		if !found {
			log.Fatalf("Registry '%s' not found", id)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Registry '%s' enabled\n", id)
	},
}

// registryDisableCmd disables a registry
var registryDisableCmd = &cobra.Command{
	Use:   "disable <id>",
	Short: "Disable a template registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		found := false
		for i, reg := range cfg.Registries {
			if reg.ID == id {
				cfg.Registries[i].Enabled = false
				found = true
				break
			}
		}

		if !found {
			log.Fatalf("Registry '%s' not found", id)
		}

		if err := config.SaveConfig(cfg); err != nil {
			log.Fatalf("Error saving configuration: %v", err)
		}

		fmt.Printf("Registry '%s' disabled\n", id)
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)

	registryCmd.AddCommand(registryListCmd)
	registryCmd.AddCommand(registryAddCmd)
	registryCmd.AddCommand(registryRemoveCmd)
	registryCmd.AddCommand(registryEnableCmd)
	registryCmd.AddCommand(registryDisableCmd)

	registryAddCmd.Flags().Bool("auth", false, "Registry requires authentication")
	registryAddCmd.Flags().String("name", "", "Human-readable name for the registry")
}
