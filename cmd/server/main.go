package main

import (
	"os"

	"github.com/rillmind/github-user-activity/src/request"
)

func main() {
	request.Request(os.Args[1])
}
