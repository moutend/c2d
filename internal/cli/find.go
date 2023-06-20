package cli

import (
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
	if len(args) < 1 || len(args[0]) == 0 {
		return nil
	}

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

	r := []rune(args[0])[0]

	cmd.Println(a.Find(r))

	return nil
}

func init() {
	RootCommand.AddCommand(findCommand)
	findCommand.PersistentFlags().StringSliceP("languages", "l", []string{}, "target languages (e.g. 'ja') (default: empty)")
}
