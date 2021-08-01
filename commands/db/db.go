package db

import (
	"github.com/spf13/cobra"
)

func NewDBCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "db",
		Short: "Manage a database from zepto project",
	}
	cmd.AddCommand(NewMigrationCmd())
	return &cmd
}
