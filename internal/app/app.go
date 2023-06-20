package app

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	HomeDir              string
	defaultDictionaryDir string
	userDictionaryDir    string
	languages            []string
	dictionary           map[rune]string
	debug                *log.Logger
}

func New() (*App, error) {
	dir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	homeDir := filepath.Join(dir, ".c2d")
	defaultDictionaryDir := filepath.Join(homeDir, "default")
	userDictionaryDir := filepath.Join(homeDir, "user")

	a := &App{
		HomeDir:              homeDir,
		defaultDictionaryDir: defaultDictionaryDir,
		userDictionaryDir:    userDictionaryDir,
		languages:            []string{},
		debug:                log.New(io.Discard, "", 0),
		dictionary:           map[rune]string{},
	}

	return a, nil
}

func (a *App) SetLanguages(languages []string) {
	a.languages = []string{}

	for i := range languages {
		a.languages = append(a.languages, strings.Join([]string{"locale", languages[i]}, string(filepath.Separator)))
	}
}
