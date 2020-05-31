package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var (
		fileFlag string
	)

	flag.StringVar(&fileFlag, "f", "", "specify file path")
	flag.Parse()

	u, err := user.Current()

	if err != nil {
		return err
	}

	rootPath := filepath.Join(u.HomeDir, ".c2d")

	userMap, err := getMap(filepath.Join(rootPath, "user.dic"))

	if err != nil {
		return err
	}

	defaultMap, err := getMap(filepath.Join(rootPath, "characterDescriptions.dic"))

	if err != nil {
		return err
	}

	if fileFlag == "" {
		return fmt.Errorf("specify file")
	}

	input, err := ioutil.ReadFile(fileFlag)

	if err != nil {
		return err
	}

	cs := []rune(string(input))

	if len(cs) < 1 {
		return nil
	}
	if desc, ok := userMap[cs[0]]; ok {
		fmt.Println(desc)

		return nil
	}
	if desc, ok := defaultMap[cs[0]]; ok {
		fmt.Println(desc)

		return nil
	}

	fmt.Println("undefined")

	return nil
}

func getMap(path string) (map[rune]string, error) {
	m := make(map[rune]string)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	ss := strings.Split(string(data), "\n")

	for _, line := range ss {
		if strings.HasPrefix(line, "#") {
			continue
		}

		rs := []rune(line)

		if len(rs) < 3 || rs[1] != rune('	') {
			continue
		}

		m[rs[0]] = string(rs[2:])
	}

	return m, nil
}
