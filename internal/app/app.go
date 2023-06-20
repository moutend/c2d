package app

import (
	"bufio"
	"io"
	"io/fs"
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

func (a *App) getDefaultDictionaryPaths(rootDir string) (paths []string, err error) {
	err = filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if !strings.HasSuffix(path, ".dic") {
			return nil
		}

		var found bool

		for i := range a.languages {
			if found = strings.Contains(path, a.languages[i]); found {
				break
			}
		}
		if found || len(a.languages) == 0 {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (a *App) getUserDictionaryPaths(rootDir string) (paths []string, err error) {
	err = filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if !strings.HasSuffix(path, ".dic") {
			return nil
		}

		paths = append(paths, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (a *App) LoadDictionaryFiles() error {
	defaultDictionaryPaths, err := a.getDefaultDictionaryPaths(a.defaultDictionaryDir)

	if err != nil {
		return err
	}

	userDictionaryPaths, err := a.getUserDictionaryPaths(a.userDictionaryDir)

	if err != nil {
		return err
	}

	paths := []string{}
	paths = append(paths, defaultDictionaryPaths...)
	paths = append(paths, userDictionaryPaths...)
	for i := range paths {
		if err := a.LoadDictionaryFile(paths[i]); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) LoadDictionaryFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), "\t")
		if len(columns) < 2 || columns[0] == "" {
			continue
		}

		r := []rune(columns[0])[0]
		description := columns[len(columns)-1]

		a.dictionary[r] = description
	}

	return nil
}