// mydocs - Simple Markdown Journal CLI
// Copyright (c) 2025 Nenad Kostic
// Licensed under the MIT License

package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// PublishHome scans for date folders and updates home.md with links to all journal entries
func PublishHome() error {
	// Find all date folders in current directory
	entries, err := findJournalEntries()
	if err != nil {
		return fmt.Errorf("failed to scan for journal entries: %v", err)
	}

	if len(entries) == 0 {
		fmt.Println(StyledInfo("No journal entries found"))
		return nil
	}

	// Read existing home.md if it exists
	existingLinks, err := readExistingHome()
	if err != nil {
		return fmt.Errorf("failed to read existing home.md: %v", err)
	}

	// Find new entries that aren't already in home.md
	newEntries := findNewEntries(entries, existingLinks)

	if len(newEntries) == 0 {
		fmt.Println(StyledInfo("home.md is already up to date"))
		return nil
	}

	// Append new entries to home.md
	return appendToHome(newEntries)
}

// JournalEntry represents a journal entry with its path and date info
type JournalEntry struct {
	Date       time.Time
	FolderName string
	FileName   string
	Link       string
}

// findJournalEntries scans current directory for folders matching YYYY-MM-DD pattern
func findJournalEntries() ([]JournalEntry, error) {
	datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	var entries []JournalEntry

	files, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		folderName := file.Name()
		if !datePattern.MatchString(folderName) {
			continue
		}

		// Parse date
		date, err := time.Parse("2006-01-02", folderName)
		if err != nil {
			continue // Skip invalid date folders
		}

		// Look for markdown file in the folder
		dayName := date.Format("Monday")
		fileName := fmt.Sprintf("%s.md", dayName)
		filePath := filepath.Join(folderName, fileName)

		if _, err := os.Stat(filePath); err == nil {
			entries = append(entries, JournalEntry{
				Date:       date,
				FolderName: folderName,
				FileName:   fileName,
				Link:       fmt.Sprintf("- [%s/%s](%s)", folderName, dayName, filePath),
			})
		}
	}

	// Sort entries by date in descending order (newest first)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Date.After(entries[j].Date)
	})

	return entries, nil
}

// readExistingHome reads existing home.md and returns existing links
func readExistingHome() (map[string]bool, error) {
	existingLinks := make(map[string]bool)

	file, err := os.Open("home.md")
	if os.IsNotExist(err) {
		return existingLinks, nil // File doesn't exist yet, that's fine
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- [") {
			existingLinks[line] = true
		}
	}

	return existingLinks, scanner.Err()
}

// findNewEntries returns entries that aren't already in home.md
func findNewEntries(entries []JournalEntry, existingLinks map[string]bool) []JournalEntry {
	var newEntries []JournalEntry

	for _, entry := range entries {
		if !existingLinks[entry.Link] {
			newEntries = append(newEntries, entry)
		}
	}

	return newEntries
}

// appendToHome appends new entries to home.md
func appendToHome(entries []JournalEntry) error {
	file, err := os.OpenFile("home.md", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open home.md: %v", err)
	}
	defer file.Close()

	for _, entry := range entries {
		if _, err := file.WriteString(entry.Link + "\n"); err != nil {
			return fmt.Errorf("failed to write to home.md: %v", err)
		}
	}

	fmt.Println(StyledSuccess(fmt.Sprintf("Added %d new entries to %s", len(entries), StyledPath("home.md"))))
	return nil
}
