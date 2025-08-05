// mydocs - Simple Markdown Journal CLI
// Copyright (c) 2025 Nenad Kostic
// Licensed under the MIT License

package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CreateEntry creates a new journal entry for the specified date.
// If date is empty, uses today's date.
func CreateEntry(date string) error {
	var targetDate time.Time
	var err error

	if date == "" {
		targetDate = time.Now()
	} else {
		targetDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return fmt.Errorf("invalid date format. Please use YYYY-MM-DD format")
		}
	}

	folderName := targetDate.Format("2006-01-02")
	dayName := targetDate.Format("Monday")
	fileName := fmt.Sprintf("%s.md", dayName)
	filePath := filepath.Join(folderName, fileName)

	// Check if folder/file already exists
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("file %s already exists. Aborting", filePath)
	}

	// Create folder if it doesn't exist
	if err := os.MkdirAll(folderName, 0755); err != nil {
		return fmt.Errorf("failed to create folder %s: %v", folderName, err)
	}

	// Create the markdown file with title
	title := fmt.Sprintf("# %s, %s", dayName, targetDate.Format("January 2, 2006"))

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	if _, err := file.WriteString(title + "\n"); err != nil {
		return fmt.Errorf("failed to write to file %s: %v", filePath, err)
	}

	fmt.Println(StyledSuccess(fmt.Sprintf("Created journal entry: %s", StyledPath(filePath))))
	return nil
}
