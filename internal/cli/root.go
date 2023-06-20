package cli

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "c2d",
	Short: "c2d -- character to description utility",
}

func init() {
	RootCommand.PersistentFlags().BoolP("debug", "d", false, "print debug messages")
}
