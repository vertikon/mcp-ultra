//go:build tools
// +build tools

package tools

import (
	_ "golang.org/x/tools/go/analysis"
	_ "golang.org/x/tools/go/analysis/singlechecker"
)
