package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/orangekame3/vercheck"
	"github.com/spf13/cobra"
)

var (
	// Version is set by goreleaser at build time
	Version   = "dev"
	repoOwner = "orangekame3"
	repoName  = "lazypoetry"
)

var noUpdate bool

var rootCmd = &cobra.Command{
	Use:   "lazypoetry",
	Short: "A lazy poetry generator and reader",
	Long: `lazypoetry is a simple CLI tool that generates random poems
and lets you browse a small collection of poetry.

Perfect for when you need a quick dose of creativity or inspiration!`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if noUpdate {
			return nil
		}
		
		// Create context with timeout for version check
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		// Check for updates using vercheck
		vercheck.CheckWithContext(ctx, vercheck.Options{
			CurrentVersion: Version,
			RepoOwner:      repoOwner,
			RepoName:       repoName,
			Silent:         true, // Don't log errors to avoid noise
		})
		
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noUpdate, "no-update", false, "disable update check")
}