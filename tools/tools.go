//go:build tools

package tools

import (
	
	_ "github.com/oligot/go-mod-upgrade"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
