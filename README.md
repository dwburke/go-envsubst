Here’s a complete, single-file README.md you can copy directly into your repo—no fragments, no extra formatting needed:

# 🧬 go-envsubst

**Go-powered envsubst with full template logic.**  
Substitute environment variables, YAML data, and Go template expressions into config files. Supports conditionals, defaults, and multiple input/output modes.

## 🚀 Features

- Go template syntax (`{{ .key }}`, `{{ env "FOO" }}`, `{{ if .debug }}`)
- Default values: `{{ env "PORT" | default "8080" }}`
- Required env vars: `{{ must_env "SECRET_KEY" }}`
- Input from:
  - Template file (`-t`)
  - Stdin (`<`)
  - Pipe (`cat file | go-envsubst`)
- Optional YAML context file (`-d`)
- Optional output file (`-o`)

## 📦 Installation

```bash
go install github.com/yourusername/go-envsubst@latest

Or build locally:

git clone https://github.com/yourusername/go-envsubst
cd go-envsubst
go build -o go-envsubst main.go

🧪 Usage

Basic substitution from env vars

export PORT=9000
go-envsubst -t config.yaml.tmpl > config.yaml

With YAML context

vars.yaml

app_name: "my-service"
debug: true

config.yaml.tmpl

app: {{ .app_name }}
port: {{ env "PORT" | default "8080" }}
debug: {{ if .debug }}true{{ else }}false{{ end }}

Command

go-envsubst -t config.yaml.tmpl -d vars.yaml -o config.yaml

Pipe + output

cat config.yaml.tmpl | go-envsubst -d vars.yaml -o config.yaml

Shell redirection

go-envsubst < config.yaml.tmpl > config.yaml

🧠 Template Functions

Function

Description

env "KEY"

Gets env var KEY, returns empty if unset

must_env "KEY"

Gets env var KEY, exits if unset

default "val"

Fallback if previous value is empty

toUpper

Converts string to uppercase

toLower

Converts string to lowercase

🛠 Example Makefile Snippet

config.yaml: config.yaml.tmpl vars.yaml
	go-envsubst -t $< -d vars.yaml -o $@

📜 License

MIT. Do whatever you want—just don’t forget to template responsibly.


Let me know if you want this wrapped into a CLI binary release, or if you’d like badges, CI setup, or a logo.

