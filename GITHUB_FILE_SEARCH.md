# GitHub File Search Tool

A command-line utility to search for specific files across GitHub repositories using the GitHub Code Search API.

## Overview

This tool allows you to search for files by name across:
- All public GitHub repositories
- Repositories owned by a specific user
- Repositories in a specific organization

## Prerequisites

- Go 1.11 or higher (for building from source)
- (Optional) GitHub Personal Access Token for higher rate limits

## Building

```bash
go build -o github-file-search github-file-search.go
```

## Usage

### Basic Syntax

```bash
./github-file-search <filename> [user/org]
```

### Arguments

- `filename` (required): The name of the file to search for
  - Exact filename: `README.md`
  - Pattern matching: `*.go`
  - Partial names work as well

- `user/org` (optional): Limit search to a specific GitHub user or organization
  - If omitted, searches across all of GitHub

### Environment Variables

- `GITHUB_TOKEN`: Your GitHub Personal Access Token
  - Highly recommended to avoid rate limiting
  - Can be created at: https://github.com/settings/tokens
  - Needs at least `public_repo` scope for public repository searches

## Examples

### Search for README.md across all GitHub
```bash
./github-file-search README.md
```

### Search for config.yml in a specific user's repositories
```bash
./github-file-search config.yml myusername
```

### Search with authentication token
```bash
GITHUB_TOKEN=ghp_xxxxxxxxxxxx ./github-file-search package.json
```

### Search for Go files in an organization
```bash
export GITHUB_TOKEN=ghp_xxxxxxxxxxxx
./github-file-search "*.go" myorg
```

## Output Format

The tool displays:
- Total number of files found
- For each result:
  - File name
  - Full path in the repository
  - Repository name
  - Direct URL to the file on GitHub

Example output:
```
Searching for file: config.yml
Limited to user/org: myusername

Found 3 file(s) matching 'config.yml'
================================================================================

[1] File: config.yml
    Path: .github/config.yml
    Repository: myusername/project1
    URL: https://github.com/myusername/project1/blob/main/.github/config.yml

[2] File: config.yml
    Path: config/config.yml
    Repository: myusername/project2
    URL: https://github.com/myusername/project2/blob/master/config/config.yml

[3] File: config.yml
    Path: config.yml
    Repository: myusername/project3
    URL: https://github.com/myusername/project3/blob/main/config.yml
```

## Rate Limits

Without authentication:
- 10 requests per minute
- 60 requests per hour

With authentication (Personal Access Token):
- 30 requests per minute
- Higher overall limits

## Error Handling

The tool provides clear error messages for common issues:
- Invalid API responses
- Network connectivity problems
- Rate limit exceeded
- Invalid authentication tokens

## Technical Details

- Uses GitHub REST API v3
- Endpoint: `https://api.github.com/search/code`
- Returns up to 100 results per search by default
- Supports GitHub's code search query syntax

## Troubleshooting

### Rate Limit Exceeded
Set up a GitHub token:
```bash
export GITHUB_TOKEN=your_token_here
```

### No Results Found
- Check spelling of the filename
- Try removing user/org restriction
- Ensure the repositories are public (private repos require additional permissions)

### Network Errors
- Check your internet connection
- Ensure you can access api.github.com
- Check if your firewall allows GitHub API access

## Use Cases

1. **Find Configuration Files**: Locate all `config.yml` files across your organization
2. **Audit Dependencies**: Search for `package.json` or `requirements.txt` files
3. **Find Documentation**: Search for README files or specific documentation
4. **Code Discovery**: Find examples of specific file types across GitHub
5. **Repository Analysis**: Identify which repositories contain certain files

## Security Notes

- Never commit your GitHub token to source control
- Use environment variables or secure secret management
- Tokens should have minimal required permissions
- Regularly rotate your access tokens

## Contributing

This is a minimal, focused tool. Future enhancements could include:
- Pagination support for more than 100 results
- Filtering by programming language
- Search by file content (not just filename)
- Export results to JSON/CSV
- Advanced query syntax support

## License

This tool is provided as-is for educational and practical use in software development workflows.
