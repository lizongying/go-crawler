package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Spider struct {
	Name string
}

//go:embed templates
var templateFS embed.FS

func main() {
	templates := map[string]string{
		"templates/cmd/example_spider/main.go.tpl":           "cmd/{n}_spider/main.go",
		"templates/spiders/example_spider/spider.go.tpl":     "internal/spiders/{n}_spider/spider.go",
		"templates/spiders/example_spider/middleware.go.tpl": "internal/spiders/{n}_spider/middleware.go",
		"templates/spiders/example_spider/types.go.tpl":      "internal/spiders/{n}_spider/types.go",
	}

	name := flag.String("n", "", "Spider name")
	force := flag.Bool("f", false, "Overwrite existing files")
	flag.Parse()

	if *name == "" {
		log.Fatal("Spider name is required, use -n <name>")
	}

	log.Println("Spider name:", *name)

	spider := Spider{Name: *name}

	for tplPath, outPattern := range templates {
		outputPath := strings.ReplaceAll(outPattern, "{n}", *name)
		if err := createFileFromTemplate(outputPath, tplPath, spider, *force); err != nil {
			log.Fatalf("Error creating %s: %v", outputPath, err)
		}
	}
}

func createFileFromTemplate(path, tplPath string, spider Spider, force bool) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if !force {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("file already exists: %s (use --force to overwrite)", path)
		}
	}

	tpl, err := template.ParseFS(templateFS, tplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if err = tpl.Execute(f, spider); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if force {
		log.Printf("Overwritten file: %s", path)
	} else {
		log.Printf("Created file: %s", path)
	}
	return nil
}
