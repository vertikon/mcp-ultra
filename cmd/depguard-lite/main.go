package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/vertikon/mcp-ultra/internal/analyzers/depguardlite"
)

func main() {
	singlechecker.Main(depguardlite.Analyzer)
}
