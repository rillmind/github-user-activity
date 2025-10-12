package main

import (
	"os"

	"github.com/rillmind/githubUserActivity/src/request"
)

func main() {
	request.Request(os.Args[1])
}
