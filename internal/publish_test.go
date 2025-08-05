package internal

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestPublishHome(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "mydocs_publish_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp directory
	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	// Create some test journal entries
	testDates := []string{"2025-01-01", "2025-01-02", "2025-01-03"}
	for _, date := range testDates {
		err := CreateEntry(date)
		if err != nil {
			t.Fatalf("Failed to create test entry for %s: %v", date, err)
		}
	}

	// Test publishing
	err = PublishHome()
	if err != nil {
		t.Errorf("PublishHome() error = %v", err)
		return
	}

	// Verify home.md was created and contains expected entries
	content, err := os.ReadFile("home.md")
	if err != nil {
		t.Errorf("Failed to read home.md: %v", err)
		return
	}

	contentStr := string(content)
	for _, date := range testDates {
		dateTime, _ := time.Parse("2006-01-02", date)
		dayName := dateTime.Format("Monday")
		expectedLink := date + "/" + dayName
		if !strings.Contains(contentStr, expectedLink) {
			t.Errorf("home.md does not contain expected link for %s", date)
		}
	}

	// Test that running publish again doesn't duplicate entries
	err = PublishHome()
	if err != nil {
		t.Errorf("Second PublishHome() error = %v", err)
		return
	}

	// Read content again after second publish
	content, err = os.ReadFile("home.md")
	if err != nil {
		t.Errorf("Failed to read home.md after second publish: %v", err)
		return
	}

	// Count occurrences of the full link line - should still be 1
	firstDate := testDates[0]
	expectedLink := "- [" + firstDate + "/"
	count := strings.Count(string(content), expectedLink)
	if count != 1 {
		t.Errorf("Expected 1 occurrence of link for %s, got %d", firstDate, count)
	}
}

func TestFindJournalEntries(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "mydocs_find_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	// Create valid journal entries
	validDates := []string{"2025-01-01", "2025-12-25"}
	for _, date := range validDates {
		CreateEntry(date)
	}

	// Create invalid folders
	os.Mkdir("invalid-folder", 0755)
	os.Mkdir("2025-13-45", 0755) // Invalid date

	entries, err := findJournalEntries()
	if err != nil {
		t.Errorf("findJournalEntries() error = %v", err)
		return
	}

	if len(entries) != len(validDates) {
		t.Errorf("Expected %d entries, got %d", len(validDates), len(entries))
	}

	// Verify entries are sorted by date (newest first)
	if len(entries) >= 2 {
		if entries[0].Date.Before(entries[1].Date) {
			t.Errorf("Entries are not sorted correctly (newest first)")
		}
	}
}

func BenchmarkPublishHome(b *testing.B) {
	tmpDir, err := os.MkdirTemp("", "mydocs_publish_bench")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	// Create many test entries
	for i := 0; i < 100; i++ {
		date := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		CreateEntry(date)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PublishHome()
	}
}
