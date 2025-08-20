// Package main is the entry point for the application.
package main

import (
	"github.com/tab-sama/oapi-codegen-custom-generator-template/internal/config"
	"github.com/tab-sama/oapi-codegen-custom-generator-template/internal/generator"
)

func main() {
	cfg := config.Load()

	gen := generator.NewGenerator(cfg)
	gen.Generate()
}
