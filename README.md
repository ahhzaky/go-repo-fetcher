# RepoFetcher

RepoFetcher is a lightweight command-line tool written in Go to fetch and clone all public repositories of a specified GitHub user. It organizes the cloned repositories into a folder named after the GitHub username.

## Features

- Fetches all public repositories of a GitHub user using the GitHub API.
- Automatically creates a folder named after the GitHub username to organize repositories.
- Clones each repository into the created folder.

## Prerequisites

- [Go](https://go.dev/) installed on your machine.
- [Git](https://git-scm.com/) installed and added to your system's PATH.

## Installation

1. Clone this repository or download the source code:
   ```bash
   git clone https://github.com/your-username/RepoFetcher.git
   cd RepoFetcher
   ```
2. Build the application:
   ```bash
   go build -o repo_fetcher
   ```

## Usage

1. Run the program from the command line:
   ```bash
   ./repo_fetcher
   ```
2. Enter the GitHub username when prompted:
   ```
   Insert username GitHub: exampleuser
   ```
3. The program will create a folder named `exampleuser` and clone all public repositories of the specified user into that folder.

## Example Output

```
Insert username GitHub: exampleuser
Cloning repository: repo1
Successfully cloned repository: repo1
Cloning repository: repo2
Successfully cloned repository: repo2
```

## Error Handling

- If the GitHub username is invalid or does not exist, the program will output an error.
- If no public repositories are found, the program will notify the user.

## Limitations

- Only public repositories are fetched. Private repositories are not accessible without authentication.
- Requires an active internet connection.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Feel free to fork the project, make changes, and submit a pull request.
