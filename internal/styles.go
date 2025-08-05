// mydocs - Simple Markdown Journal CLI
// Copyright (c) 2025 Nenad Kostic
// Licensed under the MIT License

package internal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	// Color palette
	primaryColor   = lipgloss.Color("#FF6B6B")
	secondaryColor = lipgloss.Color("#4ECDC4")
	successColor   = lipgloss.Color("#45B7D1")
	errorColor     = lipgloss.Color("#FF6B6B")
	mutedColor     = lipgloss.Color("#666666")
	accentColor    = lipgloss.Color("#96CEB4")
	highlightColor = lipgloss.Color("#FFD93D")

	// Styles
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryColor).
		MarginBottom(1)

	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(secondaryColor).
		Underline(true)

	successStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(successColor)

	errorStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(errorColor)

	infoStyle = lipgloss.NewStyle().
		Foreground(accentColor)

	commandStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(highlightColor).
		Background(lipgloss.Color("#333333")).
		Padding(0, 1)

	pathStyle = lipgloss.NewStyle().
		Foreground(mutedColor).
		Italic(true)

	helpBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(secondaryColor).
		Padding(1, 2).
		MarginTop(1)

	bannerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(highlightColor).
		Align(lipgloss.Center).
		MarginBottom(1)
)

// StyledTitle returns a styled title
func StyledTitle(text string) string {
	return titleStyle.Render(text)
}

// StyledHeader returns a styled header
func StyledHeader(text string) string {
	return headerStyle.Render(text)
}

// StyledSuccess returns a styled success message
func StyledSuccess(text string) string {
	return successStyle.Render("✅ " + text)
}

// StyledError returns a styled error message
func StyledError(text string) string {
	return errorStyle.Render("❌ " + text)
}

// StyledInfo returns a styled info message
func StyledInfo(text string) string {
	return infoStyle.Render("📝 " + text)
}

// StyledCommand returns a styled command
func StyledCommand(text string) string {
	return commandStyle.Render(text)
}

// StyledPath returns a styled path
func StyledPath(text string) string {
	return pathStyle.Render(text)
}

// StyledHelpBox returns a styled help box
func StyledHelpBox(content string) string {
	return helpBoxStyle.Render(content)
}

// StyledBanner returns a styled banner
func StyledBanner(text string) string {
	return bannerStyle.Render(text)
}

// InitLogger initializes the charm logger
func InitLogger() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(false)
}
