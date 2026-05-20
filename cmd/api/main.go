package main

import (
	"fmt"
	"time"

	"github.com/arkadashiim/go_gitify/internal/config"
	"github.com/arkadashiim/go_gitify/internal/gitwriter"
	"github.com/arkadashiim/go_gitify/internal/graph"
)

func main() {
	config, err := config.Load()
	fmt.Println(config)

	if err != nil {
		panic(err)
	}

	// gitWriter := gitwriter.NewGitWriter(config)
	// testCommit(gitWriter)

	d := graph.GraphDateByCoordinates(1, 0)
	fmt.Println(d)
}

func testCommit(gw *gitwriter.GitWriter) {
	dateString := "2025-05-19"
	date, _ := time.Parse(time.DateOnly, dateString)

	if output, err := gw.EmptyCommit(date, ""); err != nil {
		panic(err)
	} else {
		fmt.Println(string(output))
	}

	if output, err := gw.Push(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(output))
	}
}
