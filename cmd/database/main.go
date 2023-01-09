package main

import (
	"github.com/spf13/cobra"
	"github.com/xybor/xyauth/cmd/database/delete"
	"github.com/xybor/xyauth/cmd/database/migrate"
)

var rootCmd = &cobra.Command{
	Use:   "database",
	Short: "Migrate, delete, modify database",
	Long:  `A helper tool for managing database`,
}

func main() {
	rootCmd.AddCommand(
		delete.Command,
		migrate.Command,
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
