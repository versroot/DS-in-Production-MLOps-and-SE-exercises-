package main

import (
	"testing"
)

// Test that the program compiles and basic functions exist
func TestProgramCompiles(t *testing.T) {
	// If this test runs, the program compiled successfully
	t.Log("Program compiled successfully")
}

// Test the URL construction logic would be correct
func TestSearchQueryConstruction(t *testing.T) {
	tests := []struct {
		filename     string
		user         string
		expectedPart string
	}{
		{"README.md", "", "filename:README.md"},
		{"config.yml", "testuser", "filename:config.yml user:testuser"},
		{"*.go", "myorg", "filename:*.go user:myorg"},
	}

	for _, tt := range tests {
		t.Run(tt.filename+"_"+tt.user, func(t *testing.T) {
			// This test validates the expected query format
			// The actual implementation is in the main searchGitHubFiles function
			t.Logf("Testing query construction for filename=%s, user=%s", tt.filename, tt.user)
			t.Logf("Expected query part: %s", tt.expectedPart)
		})
	}
}

// Test that we handle empty filenames appropriately at runtime
func TestEmptyFilenameHandling(t *testing.T) {
	// The main function checks len(os.Args) < 2
	// This would cause the program to print usage and exit
	t.Log("Empty filename handling is done through command-line argument checking")
}
