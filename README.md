# go-envsubst

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
- **Multiple YAML context files** (`-d` flag, processed in order - later files override earlier ones)
- Optional output file (`-o`)

## Installation

```bash
go install github.com/dwburke/go-envsubst@latest
```

Or build locally:

```bash
git clone https://github.com/dwburke/go-envsubst
cd go-envsubst
go build -o go-envsubst main.go
```

## Usage

### Basic substitution from env vars

```bash
export PORT=9000
go-envsubst -t config.yaml.tmpl > config.yaml
```

### With YAML context

#### vars.yaml

```yaml
app_name: "my-service"
debug: true
```

#### config.yaml.tmpl

```yaml
app: {{ .app_name }}
port: {{ env "PORT" | default "8080" }}
debug: {{ if .debug }}true{{ else }}false{{ end }}
```

#### Command

```bash
go-envsubst -t config.yaml.tmpl -d vars.yaml -o config.yaml
```

### Multiple YAML files (layered)

Files are processed in order; later files override earlier ones:

```bash
go-envsubst -t config.yaml.tmpl -d base.yaml -d prod.yaml -o config.yaml
```

If both `base.yaml` and `prod.yaml` define `db_host`, the value from `prod.yaml` wins.

### Pipe + output

```bash
cat config.yaml.tmpl | go-envsubst -d vars.yaml -o config.yaml
```

### Shell redirection

```bash
go-envsubst < config.yaml.tmpl > config.yaml
```

## Template Functions

| Function         | Description                              |
|------------------|------------------------------------------|
| `env "KEY"`      | Gets env var KEY, returns empty if unset |
| `must_env "KEY"` | Gets env var KEY, exits if unset         |
| `default "val"`  | Fallback if previous value is empty      |
| `toUpper`        | Converts string to uppercase            |
| `toLower`        | Converts string to lowercase            |

## Example Makefile Snippet

```make
config.yaml: config.yaml.tmpl vars.yaml
    go-envsubst -t $< -d vars.yaml -o $@
```

## License

MIT License

Copyright (c) 2025 dwburke

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
