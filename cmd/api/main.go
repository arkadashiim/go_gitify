package main

import (
	"fmt"
	"time"

	"github.com/arkadashiim/go_gitify/internal/config"
	"github.com/arkadashiim/go_gitify/internal/gitwriter"
)

func main() {
	config, err := config.Load()

	if err != nil {
		panic(err)
	}

	gitWriter := gitwriter.NewGitWriter(config)

	if output, err := gitWriter.EmptyCommit(time.Now(), ""); err != nil {
		panic(err)
	} else {
		fmt.Println(string(output))
	}

	if output, err := gitWriter.Push(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(output))
	}
}
