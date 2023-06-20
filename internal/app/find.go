package app

import "fmt"

func (a *App) Find(r rune) string {
	if text, ok := a.dictionary[r]; ok {
		return text
	}

	return fmt.Sprintf("%U", r)
}
