package cmd

import (
	"fmt"
	"strconv"

	"github.com/orangekame3/lazypoetry/internal/poetry"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <id>",
	Short: "Show a specific poem by ID",
	Long:  `Display the full text of a poem from the library by its ID number.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid poem ID. Please provide a number.\n", args[0])
			return
		}

		poem, err := poetry.GetPoem(id)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			fmt.Println("\nUse 'lazypoetry list' to see available poems.")
			return
		}

		fmt.Printf("\n%s\n", poem.Title)
		fmt.Printf("by %s (%s)\n", poem.Author, poem.Form)
		fmt.Println("────────────────────")
		fmt.Printf("\n%s\n\n", poem.Content)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
