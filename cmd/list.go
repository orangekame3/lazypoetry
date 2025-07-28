package cmd

import (
	"fmt"

	"github.com/orangekame3/lazypoetry/internal/poetry"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all poems in the library",
	Long:  `Display a list of all poems available in the built-in poetry library.`,
	Run: func(cmd *cobra.Command, args []string) {
		poems := poetry.ListPoems()

		fmt.Println("\nPoetry Library:")
		fmt.Println("===============")

		for _, poem := range poems {
			fmt.Printf("ID: %-2d | %-15s | %-8s | %s\n",
				poem.ID, poem.Title, poem.Form, poem.Author)
		}

		fmt.Printf("\nTotal: %d poems\n", len(poems))
		fmt.Println("\nUse 'lazypoetry show <id>' to read a specific poem.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
