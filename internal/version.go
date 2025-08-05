// mydocs - Simple Markdown Journal CLI
// Copyright (c) 2025 Nenad Kostic
// Licensed under the MIT License

package internal

import (
	"fmt"
)

const (
	Version = "1.0.0"
	Author  = "Nenad Kostic"
	License = "MIT"
)

// ShowVersion displays version and author information
func ShowVersion() {
	versionInfo := fmt.Sprintf(`%s %s

%s
Copyright (c) 2025 %s
Licensed under the %s License

A console-based journaling utility written in Go to help you quickly 
generate daily Markdown notes and publish them to a centralized index.`,
		StyledTitle("📝 mydocs"),
		StyledCommand("v"+Version),
		StyledHeader("Author & License:"),
		Author,
		License)

	fmt.Println(StyledHelpBox(versionInfo))
}
