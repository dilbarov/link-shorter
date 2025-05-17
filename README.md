# link-shorter

This project uses [Air](https://github.com/air-verse/air) for live reloading during development.

## 📦 Development Dependencies

We use a `tools.go` file to track development tools like `air` in `go.mod`, without including them in the final build.

### File: `tools.go`

```go
//go:build tools
// +build tools

package tools

import (
	_ "github.com/air-verse/air"
)
```

## 🚀 Getting Started

1. Install Air (only needed once):

```bash
  go install github.com/air-verse/air@latest
```

Make sure `$GOPATH/bin` is in your `$PATH` or use `go env GOPATH` to find it.

2. Create .air.toml (optional) or use the default config:

```bash
  air
```

The server will restart automatically on file changes.

## ✅ Add to .gitignore

These files should not be committed:

```txt
.air
.air.toml  # if auto-generated and not customized
```

## 🧼 Clean up

To remove dev-only dependencies:

```bash
  go mod tidy
```