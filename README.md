# HeyZub 🤖

## Advanced Model Context Protocol (MCP) CLI Host

### Overview
HeyZub is a powerful Command-Line Interface (CLI) tool designed to simplify interactions with Large Language Models (LLMs) using the Model Context Protocol (MCP).

### 🌟 Key Features
- **Multi-Model Support**: Seamlessly work with different AI models
- **Dynamic Server Management**: Discover and manage MCP servers
- **Flexible Configuration**: Easily customize your MCP environment
- **Interactive Sessions**: Engage with AI models directly from the terminal

### 🚀 Installation

#### Prerequisites
- Go 1.23 or later
- Git

#### Install from Source
```bash
git clone https://github.com/ZubeidHendricks/heyZub.git
cd heyZub
go install
```

#### Install via Go
```bash
go install github.com/ZubeidHendricks/heyZub@latest
```

### 📋 Usage

#### Basic Commands
```bash
# Show version
heyzub version

# List available servers
heyzub server

# List available models
heyzub model

# Start interactive session
heyzub interact
```

### 🔧 Configuration

#### Configuration File
Create a configuration file at `~/.heyzub.yaml`:

```yaml
# Default model to use
default_model: claude-3.5-sonnet

# Configured servers
servers:
  - name: local-sqlite
    type: sqlite
    endpoint: localhost:8080
    active: true

# Available models
models:
  - name: claude-3.5-sonnet
    provider: Anthropic
    capabilities:
      - function-calling
      - context-management
```

#### Configuration Options
- `default_model`: Set the primary AI model
- `servers`: Define MCP-compatible servers
- `models`: List available AI models

### 🤝 Contributing

#### How to Contribute
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

#### Reporting Issues
- Use GitHub Issues
- Provide detailed description
- Include steps to reproduce
- Share error messages

### 📄 License
MIT License

### 🙏 Acknowledgments
- Anthropic (Claude)
- Ollama Project
- Open-source community

### 🔗 Resources
- [Model Context Protocol Specification](https://github.com/modelcontextprotocol)
- [Anthropic API Documentation](https://docs.anthropic.com)
- [Ollama Project](https://ollama.ai)
