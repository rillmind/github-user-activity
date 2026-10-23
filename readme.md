<h1 align="center"> Github Acticity CLI </h1>

<p align="center">
Simple command line interface to get github user activity
</p>

<p align="center">Project URL: https://roadmap.sh/projects/github-user-activity</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/palette/macchiato.png">
</p>

## What it does

Takes a GitHub username, calls the public endpoint `https://api.github.com/users/<username>/events/public`
(no token needed) and prints the user's recent activity, one line per event: repositories
created, stars, pushes and so on.

## Why it exists

It is the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) project from the
roadmap.sh back-end roadmap, built in Go following the project's specification: a CLI that
consumes an external API and parses its JSON, using only the standard library.

## Requirements

- Go 1.24.8 or higher
- git

## Installation

Linux:

```sh
git clone https://github.com/rillmind/githubUserActivity.git
cd githubUserActivity
go build -o github-activity ./cmd/server/
sudo cp github-activity /usr/bin/
```

Windows: 

```sh
git clone https://github.com/rillmind/githubUserActivity.git
cd githubUserActivity
go build -o github-activity.exe .\cmd\server\
```

Note: If you are on Windows, you need to add the executable binary (.exe) to the environment
variables manually. On Linux, just enter your super user password. If you want to use it without
add to environment variables, in Linux: just delete the last line and run './task-cli', in
Windows: just run normally with './task-cli' without adding or removing command lines.

## Run example

```sh
./github-activity rillmind
- Created repository rillmind/githubUserActivity
- Started catppuccin/catppuccin
- Pushed 1 commit to rillmind/taskManagerCli
- Pushed 2 commits to rillmind/taskManagerCli
- Pushed 2 commits to rillmind/taskManagerCli
- Started kamranahmedse/developer-roadmap
- Pushed 1 commit to rillmind/taskManagerCli
- Pushed 1 commit to rillmind/taskManagerCli
- Pushed 1 commit to rillmind/taskManagerCli
- Pushed 1 commit to rillmind/taskManagerCli
- Pushed 1 commit to rillmind/taskManagerCli
- Created repository rillmind/taskManagerCli
- Created repository rillmind/taskManagerCli
- Pushed 20 commits to rillmind/sauter-university-2025-challenge
- Pushed 1 commit to rillmind/dio-lab-open-source
- Pushed 1 commit to rillmind/sauter-university-2025-challenge
- Started iamgabs/sauter-university-2025-challenge
```

## How to test

There are no automated tests. `go vet ./...` checks the code; to test the behavior (needs internet):

```sh
go build -o github-activity ./cmd/server/
./github-activity rillmind                    # list of "- Created ...", "- Pushed ...", "- Started ..."
./github-activity user-that-does-not-exist-123 # json: cannot unmarshal object ... (the API returns 404)
```

Things to know when checking the output:

- Running it without a username panics (`index out of range`) instead of showing a usage message.
- Push events now show `Pushed 0 commits`: the GitHub events API stopped sending the commit list
  in the `payload`, so the count the code reads is always empty.
- Unauthenticated requests are limited to 60 per hour by GitHub.

## Code information

### Structs

Some structs, but just to get the response body in json.

```go
type GitHubEvent struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action   string   `json:"action"`
	Ref_type string   `json:"ref_type"`
	Commits  []Commit `json:"commits"`
}

type Commit struct {
	Sha string `json:"sha"`
}
```

### Functions

- **Request :** Unique function after main. Requests github api for user recently activity.

```go
func Request(username string)
```
