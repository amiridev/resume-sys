package cmd

import (
	"resume-sys/database/migrations"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Tables migration",
	Long:  `Tables will be migrate by this command`,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Migrate()
	},
}
