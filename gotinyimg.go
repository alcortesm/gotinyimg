package gotinyimg

import (
	"fmt"

	"golang.org/x/text/language"
)

func Hello(l language.Tag) (string, error) {
	switch l {
	case language.English:
		return "Hello, world!", nil
	default:
		return "", fmt.Errorf("unsupported language: %s", l)
	}
}
