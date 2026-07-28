package main

import (
	"fmt"

	"github.com/arkadashiim/go_gitify/internal/bitmap"
	"github.com/arkadashiim/go_gitify/internal/config"
	"github.com/arkadashiim/go_gitify/internal/gitwriter"
	"github.com/arkadashiim/go_gitify/internal/graph"
	"github.com/arkadashiim/go_gitify/internal/render"
)

func main() {
	config, err := config.Load()
	fmt.Println(config)

	if err != nil {
		panic(err)
	}

	gitWriter := gitwriter.NewGitWriter(config)
	testCommit(gitWriter)

	bitmapDrawer, err := bitmap.NewBitmapDrawer()
	if err != nil {
		panic(err)
	}

	bitmap, err := bitmapDrawer.DrawBitmap("печалька")
	if err != nil {
		panic(err)
	}

	render.RenderBitmap(bitmap)
	renderFullAlphabet(bitmapDrawer)
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

func renderFullAlphabet(bd *bitmap.BitmapDrawer) {
	printRune := func(r rune) {
		bitmap, err := bd.DrawBitmap(string(r))

		if err != nil {
			return
		}

		render.RenderBitmap(bitmap)
	}

	for ch := 'а'; ch <= 'я'; ch++ {
		printRune(ch)
	}
}
