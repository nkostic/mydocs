package internal

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreateEntry(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "mydocs_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp directory
	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	tests := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{"valid date", "2025-12-25", false},
		{"today's date", "", false},
		{"invalid date format", "invalid-date", true},
		{"invalid date", "2025-13-45", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateEntry(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify folder and file were created
				var expectedDate time.Time
				if tt.date == "" {
					expectedDate = time.Now()
				} else {
					expectedDate, _ = time.Parse("2006-01-02", tt.date)
				}

				folderName := expectedDate.Format("2006-01-02")
				dayName := expectedDate.Format("Monday")
				fileName := filepath.Join(folderName, dayName+".md")

				if _, err := os.Stat(fileName); os.IsNotExist(err) {
					t.Errorf("Expected file %s was not created", fileName)
				}

				// Test duplicate creation should fail
				err := CreateEntry(tt.date)
				if err == nil {
					t.Errorf("Expected error when creating duplicate entry, got nil")
				}
			}
		})
	}
}

func BenchmarkCreateEntry(b *testing.B) {
	tmpDir, err := os.MkdirTemp("", "mydocs_bench")
	if err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := time.Now().AddDate(0, 0, i).Format("2006-01-02")
		CreateEntry(date)
	}
}
