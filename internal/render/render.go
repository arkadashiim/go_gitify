package render

import (
	"fmt"

	bimapLib "github.com/arkadashiim/go_gitify/internal/bitmap"
)

func RenderBitmap(bitmap bimapLib.Bitmap) {
	for _, row := range bitmap {
		for _, point := range row {
			if point == bimapLib.Spotted {
				fmt.Print("\x1b[48;5;34m  \x1b[0m") // зелёный фон — коммит
			} else {
				fmt.Print("\x1b[48;5;236m  \x1b[0m") // серый фон — пусто
			}
		}

		fmt.Println()
	}
}
