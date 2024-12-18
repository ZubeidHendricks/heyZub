package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

const version = "0.2.0"

func main() {
    var rootCmd = &cobra.Command{
        Use:   "heyzub",
        Short: "HeyZub - Advanced MCP CLI Host",
        Long:  `Powerful CLI for Language Model interactions using Model Context Protocol`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("🤖 Welcome to HeyZub!")
            fmt.Println("Type 'heyzub help' for more information.")
        },
    }

    // Version Command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "version",
        Short: "Print HeyZub version",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("HeyZub v%s\n", version)
        },
    })

    // Server Discovery Command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "discover",
        Short: "Discover available MCP servers",
        Run: func(cmd *cobra.Command, args []string) {
            discoverServers()
        },
    })

    // Configuration Management Command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "config",
        Short: "Manage HeyZub configuration",
        Run: func(cmd *cobra.Command, args []string) {
            showConfiguration()
        },
    })

    // Model Management Command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "models",
        Short: "List and manage available models",
        Run: func(cmd *cobra.Command, args []string) {
            listModels()
        },
    })

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func discoverServers() {
    availableServers := []struct{
        Name string
        Type string
        Status string
    }{
        {"Local SQLite", "sqlite", "🟢 Active"},
        {"Filesystem Server", "filesystem", "🟢 Active"},
        {"OpenAI Server", "openai", "🔴 Inactive"},
    }

    fmt.Println("🔍 Discovered MCP Servers:")
    for _, server := range availableServers {
        fmt.Printf("- %s (%s): %s\n", server.Name, server.Type, server.Status)
    }
}

func showConfiguration() {
    viper.SetConfigName(".heyzub")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("$HOME")

    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("No configuration file found.")
        return
    }

    fmt.Println("🔧 Current HeyZub Configuration:")
    fmt.Printf("Default Model: %s\n", viper.GetString("default_model"))
    fmt.Printf("Active Servers: %v\n", viper.GetStringSlice("active_servers"))
}

func listModels() {
    models := []struct{
        Name string
        Provider string
        Capabilities []string
    }{
        {
            Name: "claude-3.5-sonnet", 
            Provider: "Anthropic", 
            Capabilities: []string{"function-calling", "context-management"},
        },
        {
            Name: "mistral-7b", 
            Provider: "Ollama", 
            Capabilities: []string{"local-inference", "multilingual"},
        },
    }

    fmt.Println("🤖 Available Models:")
    for _, model := range models {
        fmt.Printf("- %s (%s)\n", model.Name, model.Provider)
        fmt.Printf("  Capabilities: %v\n", model.Capabilities)
    }
}
