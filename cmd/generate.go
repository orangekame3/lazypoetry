package cmd

import (
	"fmt"

	"github.com/orangekame3/lazypoetry/internal/poetry"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random poem",
	Long: `Generate a random poem in the specified form.
	
Available forms:
  - haiku (default): Traditional 5-7-5 syllable Japanese poem
  - limerick: Humorous 5-line poem with AABBA rhyme scheme`,
	Run: func(cmd *cobra.Command, args []string) {
		form, _ := cmd.Flags().GetString("form")
		poem := poetry.GeneratePoem(form)
		fmt.Printf("\n%s\n\n", poem)
	},
}

func init() {
	generateCmd.Flags().StringP("form", "f", "haiku", "poem form (haiku, limerick)")
	rootCmd.AddCommand(generateCmd)
}