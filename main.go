package main

import (
    "fmt"
    "os"
    "log"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

const version = "0.3.0"

type ServerConfig struct {
    Name     string `mapstructure:"name"`
    Type     string `mapstructure:"type"`
    Endpoint string `mapstructure:"endpoint"`
    Active   bool   `mapstructure:"active"`
}

type ModelConfig struct {
    Name        string   `mapstructure:"name"`
    Provider    string   `mapstructure:"provider"`
    Capabilities []string `mapstructure:"capabilities"`
}

func main() {
    var cfgFile string

    rootCmd := &cobra.Command{
        Use:   "heyzub",
        Short: "HeyZub - Advanced MCP CLI Host",
        Long:  `Powerful CLI for Language Model interactions using Model Context Protocol`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("🤖 Welcome to HeyZub!")
            fmt.Println("Type 'heyzub help' for more information.")
        },
    }

    // Persistent Flags
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.heyzub.yaml)")

    // Version Command
    rootCmd.AddCommand(versionCmd())

    // Server Commands
    rootCmd.AddCommand(serverCmd())

    // Model Commands
    rootCmd.AddCommand(modelCmd())

    // Configuration Commands
    rootCmd.AddCommand(configCmd())

    // Interaction Command
    rootCmd.AddCommand(interactCmd())

    // Execute
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func versionCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "version",
        Short: "Print HeyZub version",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("HeyZub v%s 🚀\n", version)
            fmt.Println("Model Context Protocol CLI Host")
        },
    }
}

func serverCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "server",
        Short: "Manage MCP servers",
        Run: func(cmd *cobra.Command, args []string) {
            servers := getConfiguredServers()
            
            fmt.Println("🌐 Configured MCP Servers:")
            for _, server := range servers {
                status := "🔴 Inactive"
                if server.Active {
                    status = "🟢 Active"
                }
                fmt.Printf("- %s (%s): %s\n", server.Name, server.Type, status)
            }
        },
    }
}

func modelCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "model",
        Short: "Manage and explore AI models",
        Run: func(cmd *cobra.Command, args []string) {
            models := getConfiguredModels()
            
            fmt.Println("🤖 Available AI Models:")
            for _, model := range models {
                fmt.Printf("- %s (Provider: %s)\n", model.Name, model.Provider)
                fmt.Printf("  Capabilities: %v\n", model.Capabilities)
            }
        },
    }
}

func configCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "config",
        Short: "Manage HeyZub configuration",
        Run: func(cmd *cobra.Command, args []string) {
            initConfig()
            
            fmt.Println("🔧 HeyZub Configuration:")
            fmt.Printf("Default Model: %s\n", viper.GetString("default_model"))
            fmt.Printf("Active Servers: %v\n", viper.GetStringSlice("active_servers"))
        },
    }
}

func interactCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "interact",
        Short: "Start an interactive MCP session",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("🌟 Entering Interactive MCP Session")
            fmt.Println("Type 'exit' to quit")
            
            for {
                fmt.Print("heyzub> ")
                var input string
                fmt.Scanln(&input)
                
                if input == "exit" {
                    break
                }
                
                fmt.Printf("You entered: %s\n", input)
            }
            
            fmt.Println("Exiting interactive session. Goodbye! 👋")
        },
    }
}

func initConfig() {
    viper.SetConfigName(".heyzub")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("$HOME")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            log.Println("No configuration file found. Using defaults.")
        } else {
            log.Fatalf("Error reading config file: %s", err)
        }
    }
}

func getConfiguredServers() []ServerConfig {
    var servers []ServerConfig
    if err := viper.UnmarshalKey("servers", &servers); err != nil {
        log.Printf("Unable to decode servers config: %v", err)
        return []ServerConfig{
            {
                Name:     "Default Local Server",
                Type:     "sqlite",
                Endpoint: "localhost:8080",
                Active:   true,
            },
        }
    }
    return servers
}

func getConfiguredModels() []ModelConfig {
    var models []ModelConfig
    if err := viper.UnmarshalKey("models", &models); err != nil {
        log.Printf("Unable to decode models config: %v", err)
        return []ModelConfig{
            {
                Name:     "claude-3.5-sonnet",
                Provider: "Anthropic",
                Capabilities: []string{
                    "function-calling", 
                    "context-management", 
                    "advanced-reasoning",
                },
            },
            {
                Name:     "mistral-7b",
                Provider: "Ollama",
                Capabilities: []string{
                    "local-inference", 
                    "multilingual",
                    "open-source",
                },
            },
        }
    }
    return models
}
