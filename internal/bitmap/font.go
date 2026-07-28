package bitmap

import (
	"errors"
	"fmt"
	"strings"

	"github.com/arkadashiim/go_gitify/internal/constants"
)

type Font = map[rune][]string

func loadFont() (Font, error) {
	return validateFontMap(font)
}

func validateFontMap(fontMap Font) (Font, error) {
	var errs []error
	var errFound bool = false

	for symbol, symBitmap := range fontMap {
		symBitmapHeight := len(symBitmap)
		if symBitmapHeight > constants.DaysInWeek {
			errs = append(errs, fmt.Errorf("symbol height more than %d", constants.DaysInWeek))
			errFound = true
		}

		if errFound {
			continue
		}

		if symBitmapHeight == constants.DaysInWeek {
			continue
		}

		var symBitmapLength int = 0
		for _, bitmapRow := range symBitmap {
			if symBitmapLength > len(bitmapRow) {
				continue
			}

			symBitmapLength = len(bitmapRow)
		}

		emptyString := strings.Repeat("0", symBitmapLength)
		additionalRowsCount := constants.DaysInWeek - symBitmapHeight
		var newBitmap []string = symBitmap

		for index := range additionalRowsCount {
			if index%2 == 0 {
				newBitmap = append([]string{emptyString}, newBitmap...)
			} else {
				newBitmap = append(newBitmap, emptyString)
			}
		}

		fontMap[symbol] = newBitmap
	}

	if errFound {
		return Font{}, errors.Join(errs...)
	} else {
		return fontMap, nil
	}
}

var font = Font{
	' ': {
		"0",
	},

	'А': {
		"0110",
		"1001",
		"1111",
		"1001",
		"1001",
	},

	'Б': {
		"1111",
		"1000",
		"1111",
		"1001",
		"1111",
	},

	'В': {
		"1110",
		"1001",
		"1110",
		"1001",
		"1110",
	},

	'Г': {
		"1111",
		"1000",
		"1000",
		"1000",
		"1000",
	},

	'Д': {
		"01111",
		"01001",
		"01001",
		"11111",
		"10001",
	},

	'Е': {
		"1111",
		"1000",
		"1110",
		"1000",
		"1111",
	},

	'Ж': {
		"10101",
		"10101",
		"01110",
		"10101",
		"10101",
	},

	'З': {
		"1110",
		"0001",
		"0110",
		"0001",
		"1110",
	},

	'И': {
		"10001",
		"10011",
		"10101",
		"11001",
		"10001",
	},

	'Й': {
		"01110",
		"00000",
		"10001",
		"10011",
		"10101",
		"11001",
		"10001",
	},

	'К': {
		"1001",
		"1010",
		"1100",
		"1010",
		"1001",
	},

	'Л': {
		"0011",
		"0101",
		"1001",
		"1001",
		"1001",
	},

	'М': {
		"10001",
		"11011",
		"10101",
		"10001",
		"10001",
	},

	'Н': {
		"1001",
		"1001",
		"1111",
		"1001",
		"1001",
	},

	'О': {
		"0110",
		"1001",
		"1001",
		"1001",
		"0110",
	},

	'П': {
		"1111",
		"1001",
		"1001",
		"1001",
		"1001",
	},

	'Р': {
		"1111",
		"1001",
		"1111",
		"1000",
		"1000",
	},

	'С': {
		"0111",
		"1000",
		"1000",
		"1000",
		"0111",
	},

	'Т': {
		"11111",
		"00100",
		"00100",
		"00100",
		"00100",
	},

	'У': {
		"1001",
		"1001",
		"1111",
		"0001",
		"1111",
	},

	'Ф': {
		"00100",
		"11111",
		"10101",
		"11111",
		"00100",
	},

	'Х': {
		"10001",
		"01010",
		"00100",
		"01010",
		"10001",
	},

	'Ц': {
		"10010",
		"10010",
		"10010",
		"11111",
		"00001",
	},

	'Ч': {
		"1001",
		"1001",
		"1111",
		"0001",
		"0001",
	},

	'Ш': {
		"10101",
		"10101",
		"10101",
		"10101",
		"11111",
	},

	'Щ': {
		"101010",
		"101010",
		"101010",
		"111111",
		"000001",
	},

	'Ъ': {
		"1100",
		"0100",
		"0111",
		"0101",
		"0111",
	},

	'Ы': {
		"10001",
		"10001",
		"11101",
		"10101",
		"11101",
	},

	'Ь': {
		"100",
		"100",
		"111",
		"101",
		"111",
	},

	'Э': {
		"1110",
		"0001",
		"1111",
		"0001",
		"1110",
	},

	'Ю': {
		"10010",
		"10101",
		"11101",
		"10101",
		"10010",
	},

	'Я': {
		"1111",
		"1001",
		"1111",
		"0101",
		"1001",
	},
}
