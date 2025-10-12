<h1 align="center"> Github Acticity CLI </h1>

<p align="center">
Simple command line interface to get github user activity
</p>

<p align="center">Project URL: https://roadmap.sh/projects/github-user-activity</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/palette/macchiato.png">
</p>

## Requirements

- Go 1.24.8 or higher
- git

## Installation

Linux:

```sh
git clone https://github.com/rillmind/githubUserActivity.git
cd githubUserActivity
go build -o github-activity ./cmd
sudo cp github-activity /usr/bin/
```

Windows: 

```sh
git clone https://github.com/rillmind/githubUserActivity.git
cd githubUserActivity
go build -o github-activity.exe .\cmd\
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
