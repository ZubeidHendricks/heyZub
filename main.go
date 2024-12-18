package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "heyzub",
        Short: "HeyZub - Advanced MCP CLI Host",
        Long:  `Powerful CLI for language model interactions`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("🤖 Welcome to HeyZub!")
            fmt.Println("Type 'heyzub help' for more information.")
        },
    }

    // Add version command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "version",
        Short: "Print version information",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("HeyZub v0.1.0")
        },
    })

    // Add server command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "server",
        Short: "Manage MCP servers",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Server Management")
            fmt.Println("Available servers:")
            fmt.Println("- Local SQLite")
            fmt.Println("- Filesystem")
        },
    })

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
