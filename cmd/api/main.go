package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/arkadashiim/go_gitify/internal/bitmap"
	"github.com/arkadashiim/go_gitify/internal/config"
	"github.com/arkadashiim/go_gitify/internal/constants"
	"github.com/arkadashiim/go_gitify/internal/gitwriter"
	"github.com/arkadashiim/go_gitify/internal/graph"
	"github.com/arkadashiim/go_gitify/internal/render"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) > 1 && os.Args[1] == "--reset" {
		resetGraph(reader)
		return
	}

	text := textFromArgsOrPrompt(reader)
	start := graph.StartDate(yearFromPrompt(reader))

	bitmapDrawer, err := bitmap.NewBitmapDrawer()
	if err != nil {
		panic(err)
	}

	bm, err := bitmapDrawer.DrawBitmap(text)
	if err != nil {
		panic(err)
	}

	render.RenderBitmap(bm, start)

	if !confirm(reader, "Записать это в git? [Y/n] ") {
		fmt.Println("Отменено")
		return
	}

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	gitWriter := gitwriter.NewGitWriter(cfg)

	if err := fillGraph(gitWriter, bm, start); err != nil {
		panic(err)
	}

	output, err := gitWriter.Push()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

func textFromArgsOrPrompt(reader *bufio.Reader) string {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " ")
	}

	fmt.Print("Текст для графика: ")
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}

func yearFromPrompt(reader *bufio.Reader) int {
	fmt.Print("Год графика (Enter — текущий, либо год, например 2023): ")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			return 0
		}

		year, err := strconv.Atoi(input)
		if err == nil && year > 0 {
			return year
		}

		fmt.Print("Не похоже на год, попробуй ещё раз: ")
	}
}

func fillGraph(gw *gitwriter.GitWriter, bm bitmap.Bitmap, start time.Time) error {
	for day := range constants.DaysInWeek {
		for week := range constants.WeeksInYear {
			if bm[day][week] != bitmap.Spotted {
				continue
			}

			date, err := graph.GraphDateByCoordinates(start, week, day)
			if err != nil {
				return err
			}

			if _, err := gw.EmptyCommit(date, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func resetGraph(reader *bufio.Reader) {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Это сотрёт ВСЮ историю коммитов в %s (ветка %s) и форс-запушит пустое состояние в origin. Это необратимо.\n",
		cfg.RepoPath, cfg.GraphBranch,
	)
	fmt.Print("Чтобы подтвердить, введи RESET: ")

	input, _ := reader.ReadString('\n')
	if strings.TrimSpace(input) != "RESET" {
		fmt.Println("Отменено")
		return
	}

	gitWriter := gitwriter.NewGitWriter(cfg)

	if output, err := gitWriter.ResetHistory(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(output))
	}

	if output, err := gitWriter.ForcePush(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(output))
	}
}

func confirm(reader *bufio.Reader, prompt string) bool {
	fmt.Print(prompt)

	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))

	return answer == "" || answer == "y" || answer == "yes" || answer == "д" || answer == "да"
}
