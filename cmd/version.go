package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Display the current version of lazypoetry.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("lazypoetry %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}