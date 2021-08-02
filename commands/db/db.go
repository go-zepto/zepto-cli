package db

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/go-zepto/zepto/config"
	"github.com/go-zepto/zepto/database/dbtools"
	"github.com/spf13/cobra"
)

func NewDBCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "db",
		Short: "Manage a database from zepto project",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create the project database",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := config.NewConfigFromFile()
			dt, err := dbtools.NewDBTools()
			if err != nil {
				return err
			}
			err = dt.CreateDB()
			if err != nil {
				return err
			}
			action := color.New(color.FgGreen, color.Bold).Sprint("created")
			fmt.Printf("Database %s %s successfully.\n", c.DB.Database, action)
			return nil
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "drop",
		Short: "Drop the project database",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := config.NewConfigFromFile()
			dt, err := dbtools.NewDBTools()
			if err != nil {
				return err
			}
			err = dt.DropDB()
			if err != nil {
				return err
			}
			action := color.New(color.FgYellow, color.Bold).Sprint("dropped")
			fmt.Printf("Database %s %s successfully.\n", c.DB.Database, action)
			return nil
		},
	})
	cmd.AddCommand(NewMigrationCmd())
	return &cmd
}
