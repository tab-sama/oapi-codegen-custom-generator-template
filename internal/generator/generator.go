// Package generator provides functionality for generating code from OpenAPI specifications.
package generator

import (
	"fmt"
	"log"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/util"

	"github.com/tab-sama/oapi-codegen-custom-generator-template/internal/config"
)

// Generator handles the code generation process based on OpenAPI specification files.
type Generator struct {
	filePath string
}

// NewGenerator creates a new Generator instance with the provided configuration.
func NewGenerator(cfg *config.Config) *Generator {
	return &Generator{
		filePath: cfg.FilePath,
	}
}

// Generate loads the OpenAPI specification and performs the code generation process.
func (g *Generator) Generate() {
	swagger, err := util.LoadSwagger(g.filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load spec: %v", err))
	}

	log.Println(swagger)
}
