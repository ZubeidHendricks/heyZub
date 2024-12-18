package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err = root.Execute()
	output = buf.String()

	return output, err
}

func TestVersionCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "heyzub"}
	rootCmd.AddCommand(versionCmd())

	output, err := executeCommand(rootCmd, "version")
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedOutputs := []string{
		fmt.Sprintf("HeyZub v%s", version),
		"Model Context Protocol CLI Host",
	}

	for _, expected := range expectedOutputs {
		if !bytes.Contains([]byte(output), []byte(expected)) {
			t.Errorf("Expected output to contain '%s', but got: %s", expected, output)
		}
	}
}

func TestServerCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "heyzub"}
	rootCmd.AddCommand(serverCmd())

	output, err := executeCommand(rootCmd, "server")
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedOutputs := []string{
		"Configured MCP Servers:",
		"Default Local Server",
		"sqlite",
	}

	for _, expected := range expectedOutputs {
		if !bytes.Contains([]byte(output), []byte(expected)) {
			t.Errorf("Expected output to contain '%s', but got: %s", expected, output)
		}
	}
}

func TestModelCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "heyzub"}
	rootCmd.AddCommand(modelCmd())

	output, err := executeCommand(rootCmd, "model")
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedModels := []string{
		"claude-3.5-sonnet",
		"mistral-7b",
		"Anthropic",
		"Ollama",
	}

	for _, expected := range expectedModels {
		if !bytes.Contains([]byte(output), []byte(expected)) {
			t.Errorf("Expected output to contain '%s', but got: %s", expected, output)
		}
	}
}

func TestConfigCommand(t *testing.T) {
	// Create a temporary config file
	tmpfile, err := os.CreateTemp("", "heyzub-test-config-*.yaml")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Write sample config
	_, err = tmpfile.Write([]byte(`
default_model: test-model
active_servers:
  - test-server
`))
	if err != nil {
		t.Fatalf("Cannot write to temporary file: %v", err)
	}
	tmpfile.Close()

	// Set up Viper to use the temp config file
	os.Setenv("HOME", "/tmp")

	rootCmd := &cobra.Command{Use: "heyzub"}
	rootCmd.AddCommand(configCmd())

	output, err := executeCommand(rootCmd, "config")
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedOutputs := []string{
		"HeyZub Configuration:",
		"Default Model:",
		"Active Servers:",
	}

	for _, expected := range expectedOutputs {
		if !bytes.Contains([]byte(output), []byte(expected)) {
			t.Errorf("Expected output to contain '%s', but got: %s", expected, output)
		}
	}
}

func TestInteractCommand(t *testing.T) {
	rootCmd := &cobra.Command{Use: "heyzub"}
	rootCmd.AddCommand(interactCmd())

	// Simulate user input and exit
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte("exit\n"))
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	output, err := executeCommand(rootCmd, "interact")
	
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedOutputs := []string{
		"Entering Interactive MCP Session",
		"Type 'exit' to quit",
		"Exiting interactive session",
	}

	for _, expected := range expectedOutputs {
		if !bytes.Contains([]byte(output), []byte(expected)) {
			t.Errorf("Expected output to contain '%s', but got: %s", expected, output)
		}
	}
}

func TestConfiguredServers(t *testing.T) {
	servers := getConfiguredServers()

	if len(servers) == 0 {
		t.Error("Expected at least one configured server")
	}

	expectedServer := ServerConfig{
		Name:     "Default Local Server",
		Type:     "sqlite",
		Endpoint: "localhost:8080",
		Active:   true,
	}

	found := false
	for _, server := range servers {
		if server.Name == expectedServer.Name && 
		   server.Type == expectedServer.Type && 
		   server.Endpoint == expectedServer.Endpoint && 
		   server.Active == expectedServer.Active {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Could not find expected default server configuration")
	}
}

func TestConfiguredModels(t *testing.T) {
	models := getConfiguredModels()

	if len(models) == 0 {
		t.Error("Expected at least one configured model")
	}

	expectedModels := []string{"claude-3.5-sonnet", "mistral-7b"}
	
	for _, expectedName := range expectedModels {
		found := false
		for _, model := range models {
			if model.Name == expectedName {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Could not find expected model: %s", expectedName)
		}
	}
}
