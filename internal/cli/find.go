package cli

import (
	"os"

	"github.com/moutend/c2d/internal/app"
	"github.com/spf13/cobra"
)

var findCommand = &cobra.Command{
	Use:     "find",
	Aliases: []string{"f"},
	Short:   "find the description of the given character",
	RunE:    findCommandRunE,
}

func findCommandRunE(cmd *cobra.Command, args []string) error {
	a, err := app.New()

	if err != nil {
		return err
	}

	languages, err := cmd.Flags().GetStringSlice("languages")

	if len(languages) > 0 {
		a.SetLanguages(languages)
	}
	if err := a.LoadDictionaryFiles(); err != nil {
		return err
	}

	filePath, err := cmd.Flags().GetString("file")

	if err != nil {
		return err
	}

	var r rune

	if filePath != "" {
		data, err := os.ReadFile(filePath)

		if err != nil {
			return err
		}
		if len(data) < 1 {
			return nil
		}

		r = []rune(string(data))[0]
	} else {
		r = []rune(args[0])[0]
	}

	cmd.Println(a.Find(r))

	return nil
}

func init() {
	RootCommand.AddCommand(findCommand)

	findCommand.PersistentFlags().StringP("file", "f", "", "read from file")
	findCommand.PersistentFlags().StringSliceP("languages", "l", []string{}, "target languages (e.g. 'ja') (default: empty)")
}
