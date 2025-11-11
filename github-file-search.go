package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// GitHubSearchResponse represents the GitHub code search API response
type GitHubSearchResponse struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		HTMLURL    string `json:"html_url"`
		Repository struct {
			FullName string `json:"full_name"`
			HTMLURL  string `json:"html_url"`
		} `json:"repository"`
	} `json:"items"`
}

func searchGitHubFiles(filename string, user string, token string) error {
	// Build the search query
	query := fmt.Sprintf("filename:%s", filename)
	if user != "" {
		query += fmt.Sprintf(" user:%s", user)
	}

	// Construct the API URL
	apiURL := fmt.Sprintf("https://api.github.com/search/code?q=%s", url.QueryEscape(query))

	// Create HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Add headers
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse the JSON response
	var searchResult GitHubSearchResponse
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	// Display results
	fmt.Printf("\nFound %d file(s) matching '%s'\n", searchResult.TotalCount, filename)
	fmt.Println(strings.Repeat("=", 80))

	for i, item := range searchResult.Items {
		fmt.Printf("\n[%d] File: %s\n", i+1, item.Name)
		fmt.Printf("    Path: %s\n", item.Path)
		fmt.Printf("    Repository: %s\n", item.Repository.FullName)
		fmt.Printf("    URL: %s\n", item.HTMLURL)
	}

	if searchResult.TotalCount == 0 {
		fmt.Println("\nNo files found matching your search criteria.")
	}

	return nil
}

func printUsage() {
	fmt.Println("GitHub File Search Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  github-file-search <filename> [user/org]")
	fmt.Println("\nArguments:")
	fmt.Println("  filename    : The name of the file to search for (e.g., 'config.yml', '*.go')")
	fmt.Println("  user/org    : (Optional) Limit search to a specific user or organization")
	fmt.Println("\nEnvironment Variables:")
	fmt.Println("  GITHUB_TOKEN: GitHub personal access token for authentication (recommended)")
	fmt.Println("                Without a token, you may hit rate limits quickly.")
	fmt.Println("\nExamples:")
	fmt.Println("  github-file-search README.md")
	fmt.Println("  github-file-search config.yml myusername")
	fmt.Println("  GITHUB_TOKEN=ghp_xxx github-file-search package.json myorg")
	fmt.Println("\nNote: To search across all of GitHub, use without the user/org parameter.")
	fmt.Println("      For better results, consider providing a GitHub token.")
}

func main() {
	// Check arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse arguments
	filename := os.Args[1]
	user := ""
	if len(os.Args) >= 3 {
		user = os.Args[2]
	}

	// Get GitHub token from environment
	token := os.Getenv("GITHUB_TOKEN")

	// Perform the search
	fmt.Printf("Searching for file: %s\n", filename)
	if user != "" {
		fmt.Printf("Limited to user/org: %s\n", user)
	} else {
		fmt.Println("Searching across all GitHub repositories")
	}

	if token == "" {
		fmt.Println("\nWarning: No GITHUB_TOKEN found. You may encounter rate limits.")
		fmt.Println("Set GITHUB_TOKEN environment variable for better experience.\n")
	}

	err := searchGitHubFiles(filename, user, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
