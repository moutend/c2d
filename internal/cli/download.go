package cli

import (
	"fmt"
	"net/url"

	"github.com/moutend/c2d/internal/app"
	"github.com/spf13/cobra"
)

var downloadCommand = &cobra.Command{
	Use:     "download",
	Aliases: []string{"d"},
	Short:   "download download files",
	RunE:    downloadCommandRunE,
}

func downloadCommandRunE(cmd *cobra.Command, args []string) error {
	a, err := app.New()

	if err != nil {
		return err
	}

	repo, err := cmd.Flags().GetString("repo")

	if err != nil {
		return err
	}

	repoURL, err := url.Parse(repo)

	if err != nil {
		return err
	}

	paths, err := cmd.Flags().GetStringSlice("paths")

	if err != nil {
		return err
	}
	if err := a.DownloadDictionaryFiles(repoURL, paths...); err != nil {
		return err
	}

	return nil
}

func init() {
	RootCommand.AddCommand(downloadCommand)

	downloadCommand.PersistentFlags().StringP("repo", "r", app.RepoURL, fmt.Sprintf("repository URL (default: %q", app.RepoURL))
	downloadCommand.PersistentFlags().StringSliceP("paths", "p", app.DictionaryFilePaths, "file paths")
}
