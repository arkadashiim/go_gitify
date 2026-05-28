package main

import (
	"fmt"

	"github.com/arkadashiim/go_gitify/internal/bitmap"
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

	gitWriter := gitwriter.NewGitWriter(config)
	testCommit(gitWriter)

	font, err := bitmap.LoadFont()
	if err != nil {
		panic(err)
	}

	fmt.Println(font[' '])
}

func testCommit(gw *gitwriter.GitWriter) {
	return

	date, err := graph.GraphDateByCoordinates(1, 0)
	if err != nil {
		panic(err)
	}

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

	fmt.Println(date)
}
