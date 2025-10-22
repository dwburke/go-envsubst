package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func main() {
	var tmplPath string
	var dataPaths multiStringFlag
	var outPath string

	flag.StringVar(&tmplPath, "t", "", "Template file path (optional if using stdin)")
	flag.Var(&dataPaths, "d", "YAML data file paths (can specify multiple)")
	flag.StringVar(&outPath, "o", "", "Output file path (optional)")
	flag.Parse()

	_ = godotenv.Load()

	// Load and merge YAML data if provided
	// Files are processed in the order they are specified on the command line.
	// Variables from later files will overwrite those from earlier ones.
	data := map[string]any{}
	for _, dataPath := range dataPaths {
		yamlBytes, err := os.ReadFile(dataPath)
		if err != nil {
			log.Fatalf("Error reading YAML file %s: %v", dataPath, err)
		}
		tempData := map[string]any{}
		if err := yaml.Unmarshal(yamlBytes, &tempData); err != nil {
			log.Fatalf("Error parsing YAML file %s: %v", dataPath, err)
		}
		// Merge variables, allowing later files to overwrite earlier ones
		for k, v := range tempData {
			data[k] = v
		}
	}

	funcs := template.FuncMap{
		"env": func(key string) string {
			return os.Getenv(key)
		},
		"must_env": func(key string) string {
			val := os.Getenv(key)
			if val == "" {
				log.Fatalf("Missing required env var: %s", key)
			}
			return val
		},
		"default": func(def string) func(string) string {
			return func(val string) string {
				if val == "" {
					return def
				}
				return val
			}
		},
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
	}

	var tmplReader io.Reader
	if tmplPath != "" {
		file, err := os.Open(tmplPath)
		if err != nil {
			log.Fatalf("Error opening template file: %v", err)
		}
		defer file.Close()
		tmplReader = file
	} else {
		tmplReader = os.Stdin
	}

	tmplBytes, err := io.ReadAll(tmplReader)
	if err != nil {
		log.Fatalf("Error reading template: %v", err)
	}

	tmpl, err := template.New("tmpl").Funcs(funcs).Parse(string(tmplBytes))
	if err != nil {
		log.Fatalf("Template parse error: %v", err)
	}

	var output io.Writer
	if outPath != "" {
		file, err := os.Create(outPath)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	if err := tmpl.Execute(output, data); err != nil {
		log.Fatalf("Template execution error: %v", err)
	}
}

type multiStringFlag []string

func (m *multiStringFlag) String() string {
	return strings.Join(*m, ",")
}

func (m *multiStringFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}
