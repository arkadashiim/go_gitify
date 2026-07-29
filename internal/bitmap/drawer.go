package bitmap

import (
	"fmt"
	"unicode"

	"github.com/arkadashiim/go_gitify/internal/constants"
)

type BitmapPoint = rune
type Bitmap = [constants.DaysInWeek][constants.WeeksInYear]BitmapPoint

const (
	Empty   BitmapPoint = 0
	Spotted BitmapPoint = 1
)

const DefaultBitmapPoint BitmapPoint = Empty

const LetterGap = 1

type BitmapDrawer struct {
	font Font
}

func NewBitmapDrawer() (*BitmapDrawer, error) {
	font, err := loadFont()

	if err != nil {
		return &BitmapDrawer{}, fmt.Errorf("failed init bitmap drawer: %w", err)
	}

	return &BitmapDrawer{font: font}, nil
}

func (bd *BitmapDrawer) DrawBitmap(text string) (Bitmap, error) {
	var bitmap Bitmap
	totalWidth := 0

	// подсчет длины заранее для центрирования
	for _, char := range text {
		word, ok := font[unicode.ToUpper(char)]

		if !ok {
			return Bitmap{}, fmt.Errorf("there is no char \"%s\" in font", string(char))
		}

		maxRowLength := 0
		for _, symbolsRow := range word {
			if len(symbolsRow) <= maxRowLength {
				continue
			}

			maxRowLength = len(symbolsRow)
		}

		totalWidth += maxRowLength + LetterGap
	}

	// после последней буквы зазор не нужен
	totalWidth -= LetterGap

	if totalWidth > constants.WeeksInYear {
		return Bitmap{}, fmt.Errorf("bitmap length exceeded!")
	}

	offset := (constants.WeeksInYear - totalWidth) / 2
	for _, char := range text {
		word, ok := font[unicode.ToUpper(char)]

		if !ok {
			return Bitmap{}, fmt.Errorf("there is no char \"%s\" in font", string(char))
		}

		maxRowLength := 0
		for _, symbolsRow := range word {
			if len(symbolsRow) <= maxRowLength {
				continue
			}

			maxRowLength = len(symbolsRow)
		}

		for rowIndex, symbolsRow := range word {

			for symbolIndex, symbol := range symbolsRow {
				currentSymbolIndex := offset + symbolIndex

				if symbol == '1' {
					bitmap[rowIndex][currentSymbolIndex] = Spotted
				} else {
					bitmap[rowIndex][currentSymbolIndex] = Empty
				}
			}
		}

		offset += maxRowLength + LetterGap
	}

	return bitmap, nil
}
