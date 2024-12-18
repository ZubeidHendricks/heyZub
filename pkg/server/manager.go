package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// ServerType represents different types of MCP servers
type ServerType string

const (
	SQLiteServer     ServerType = "sqlite"
	FileSystemServer ServerType = "filesystem"
	OpenAIServer     ServerType = "openai"
)

// ServerConfig represents the configuration for an MCP server
type ServerConfig struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Type     ServerType `json:"type"`
	Endpoint string     `json:"endpoint"`
	Active   bool       `json:"active"`
	Config   string     `json:"config,omitempty"`
}

// ServerManager handles discovery, registration, and management of MCP servers
type ServerManager struct {
	servers     map[string]ServerConfig
	mutex       sync.RWMutex
	configPath  string
}

// NewServerManager creates a new server manager
func NewServerManager() *ServerManager {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Could not determine user config directory:", err)
	}

	heyZubDir := filepath.Join(configDir, "heyzub")
	err = os.MkdirAll(heyZubDir, 0755)
	if err != nil {
		log.Fatal("Could not create heyzub config directory:", err)
	}

	return &ServerManager{
		servers:    make(map[string]ServerConfig),
		configPath: filepath.Join(heyZubDir, "servers.json"),
	}
}

// RegisterServer adds a new server to the manager
func (sm *ServerManager) RegisterServer(server ServerConfig) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// Validate server configuration
	if server.Name == "" || server.Endpoint == "" {
		return fmt.Errorf("server name and endpoint are required")
	}

	// Generate unique ID if not provided
	if server.ID == "" {
		server.ID = generateUniqueID()
	}

	sm.servers[server.ID] = server
	return sm.saveServers()
}

// UnregisterServer removes a server from the manager
func (sm *ServerManager) UnregisterServer(serverID string) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if _, exists := sm.servers[serverID]; !exists {
		return fmt.Errorf("server not found")
	}

	delete(sm.servers, serverID)
	return sm.saveServers()
}

// ListServers returns all registered servers
func (sm *ServerManager) ListServers() []ServerConfig {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	servers := make([]ServerConfig, 0, len(sm.servers))
	for _, server := range sm.servers {
		servers = append(servers, server)
	}
	return servers
}

// LoadServers reads servers from configuration file
func (sm *ServerManager) LoadServers() error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	data, err := os.ReadFile(sm.configPath)
	if os.IsNotExist(err) {
		// Initialize with default servers if no config exists
		sm.initDefaultServers()
		return sm.saveServers()
	} else if err != nil {
		return fmt.Errorf("error reading server configuration: %v", err)
	}

	return json.Unmarshal(data, &sm.servers)
}

// saveServers writes servers to configuration file
func (sm *ServerManager) saveServers() error {
	data, err := json.MarshalIndent(sm.servers, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling servers: %v", err)
	}

	return os.WriteFile(sm.configPath, data, 0644)
}

// initDefaultServers populates with default server configurations
func (sm *ServerManager) initDefaultServers() {
	defaultServers := []ServerConfig{
		{
			ID:       "local-sqlite",
			Name:     "Local SQLite Server",
			Type:     SQLiteServer,
			Endpoint: "localhost:8080",
			Active:   true,
		},
		{
			ID:       "local-filesystem",
			Name:     "Local Filesystem Server",
			Type:     FileSystemServer,
			Endpoint: "/tmp/mcp-server",
			Active:   true,
		},
	}

	for _, server := range defaultServers {
		sm.servers[server.ID] = server
	}
}

// generateUniqueID creates a unique identifier for servers
func generateUniqueID() string {
	// In a real-world scenario, use a more robust ID generation method
	return fmt.Sprintf("server-%d", os.Getpid())
}

// ValidateServerConfig checks if a server configuration is valid
func (sm *ServerManager) ValidateServerConfig(server ServerConfig) error {
	if server.Name == "" {
		return fmt.Errorf("server name cannot be empty")
	}

	switch server.Type {
	case SQLiteServer:
		// Add specific SQLite validation logic
	case FileSystemServer:
		// Add specific filesystem validation logic
	case OpenAIServer:
		// Add specific OpenAI server validation logic
	default:
		return fmt.Errorf("unsupported server type: %s", server.Type)
	}

	return nil
}
