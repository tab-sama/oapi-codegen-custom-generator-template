# 🚀 Custom oapi-codegen Generator Template

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> **Note**: This is a GitHub template repository that demonstrates how to create custom code generators using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) with Go templates.

## 📋 About This Template

This project shows how to build custom code generators that process OpenAPI specifications using oapi-codegen as a foundation and Go templates to define the generated code output. Instead of generating standard client/server boilerplate, you can create completely custom code tailored to your specific needs.

## 🎯 What This Template Demonstrates

- **Custom Code Generation**: Generate any type of code from OpenAPI specs using Go templates
- **Template-Based Approach**: Use Go's `text/template` package to define exactly what code to generate
- **Flexible Configuration**: Easy configuration through environment variables
- **Real Example**: Working example that generates Go functions from a User Management API

## 🏗️ Project Structure

```
├── cmd/generate/main.go          # Entry point for the code generator
├── internal/
│   ├── config/config.go          # Configuration management
│   └── generator/generator.go    # Core generation logic
├── templates/
│   └── functions.go.tmpl         # Go template defining the output format
├── api/
│   └── openapi.yml              # Sample OpenAPI specification
└── README.md                    # This file
```

## 🔧 How It Works

1. **Load OpenAPI Spec**: Uses `oapi-codegen/pkg/util.LoadSwagger()` to parse the OpenAPI specification
2. **Extract Data**: Processes the spec to extract operations (GET, POST, PUT, DELETE) and their details
3. **Apply Templates**: Uses Go templates to generate custom code based on the extracted data
4. **Output Generation**: Produces the final code according to your template definitions

### Example Output

Given the sample OpenAPI spec, the generator creates Go functions like:

```go
// Generated functions from OpenAPI specification

package main

import "log"

// listUsers handles GET /users
func listUsers() {
	log.Println("read only")
}

// createUser handles POST /users
func createUser() {
	log.Println("write only")
}

// getUserByID handles GET /users/{id}
func getUserByID() {
	log.Println("read only")
}

// updateUser handles PUT /users/{id}
func updateUser() {
	log.Println("write only")
}

// deleteUser handles DELETE /users/{id}
func deleteUser() {
	log.Println("write only")
}
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.5 or later
- (Optional) [Moonrepo](https://moonrepo.dev/) for enhanced build tooling

### Setup

1. **Use this template** by clicking "Use this template" button on GitHub
2. **Clone your new repository**:
   ```bash
   git clone <your-repo-url>
   cd <your-repo-name>
   ```

3. **Install dependencies**:
   ```bash
   # Using Moonrepo (recommended)
   proto use
   moon :setup

   # Using Make
   make install

   # Using Go directly
   go mod download
   ```

### Running the Generator

```bash
# Using Moonrepo
moon :run

# Using Make
make run

# Using Go directly
go run ./cmd/generate

# With custom OpenAPI spec
FILE_PATH=./path/to/your/openapi.yml go run ./cmd/generate
```

## 🛠️ Customization

### 1. Modify the Template

Edit `templates/functions.go.tmpl` to change the generated code format. The template receives an array of `OperationData` structs:

```go
type OperationData struct {
    OperationID string  // e.g., "listUsers"
    Method      string  // e.g., "GET"
    Path        string  // e.g., "/users"
    LogMessage  string  // e.g., "read only"
}
```

### 2. Extend the Generator

Modify `internal/generator/generator.go` to:
- Extract additional data from the OpenAPI spec
- Process different parts of the specification (schemas, parameters, etc.)
- Add custom logic for your specific use case

### 3. Add More Templates

Create additional template files and modify the generator to use them:

```go
// Example: Add multiple templates
clientTmpl, _ := template.ParseFiles("templates/client.go.tmpl")
serverTmpl, _ := template.ParseFiles("templates/server.go.tmpl")
```

### 4. Configuration Options

Add more configuration options in `internal/config/config.go`:

```go
type Config struct {
    FilePath     string `default:"./api/openapi.yml" envconfig:"FILE_PATH"`
    OutputDir    string `default:"./generated" envconfig:"OUTPUT_DIR"`
    PackageName  string `default:"generated" envconfig:"PACKAGE_NAME"`
}
```

## 📝 Example Use Cases

- **API Client Generators**: Generate custom HTTP clients
- **Mock Generators**: Create mock implementations for testing
- **Documentation**: Generate custom API documentation
- **Validation Code**: Create request/response validators
- **Database Models**: Generate database models from schemas
- **Configuration Files**: Generate config files for other tools

## 🔨 Available Commands

### Moonrepo Commands

- `moon :setup` - Set up project dependencies
- `moon :run` - Run the code generator
- `moon :build` - Build the project
- `moon :test` - Run tests
- `moon :lint` - Lint the code
- `moon :format` - Format the code
- `moon :clean` - Clean build artifacts

### Make Commands

- `make install` - Install tools and dependencies
- `make run` - Run the code generator
- `make build` - Build the project
- `make test` - Run tests
- `make lint` - Lint the code
- `make format` - Format the code
- `make clean` - Clean build artifacts

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📚 Resources

- [oapi-codegen Documentation](https://github.com/oapi-codegen/oapi-codegen)
- [Go Templates Documentation](https://pkg.go.dev/text/template)
- [OpenAPI Specification](https://swagger.io/specification/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
