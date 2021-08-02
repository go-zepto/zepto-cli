package db

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/go-zepto/zepto/database/migrate"
	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"
)

func handleRunErr(err error) error {
	if errors.Is(err, gomigrate.ErrNoChange) {
		color.Yellow("Migration was executed, but there were no changes.")
		return nil
	}
	if errors.Is(err, gomigrate.ErrNilVersion) {
		color.Yellow("No migrations have been run yet. Please run:")
		fmt.Printf("\n  $ zepto db migration up\n\n")
		return nil
	}
	color.Green("Migration executed successfully")
	return err
}

func NewMigrationCmd() *cobra.Command {
	migrationCmd := cobra.Command{
		Use:   "migration",
		Short: "Manage a database migrations",
	}
	migrationCmd.AddCommand(&cobra.Command{
		Use:   "create [name]",
		Short: "Create migration files (up and down)",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := migrate.NewMigrate(migrate.Options{})
			if err != nil {
				return err
			}
			return m.CreateMigrationFiles(migrate.CreateMigrationFilesOptions{
				Name: args[0],
			})
		},
	})
	migrationCmd.AddCommand(&cobra.Command{
		Use:   "up [steps]",
		Short: "Run migration up",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := migrate.NewMigrate(migrate.Options{})
			if err != nil {
				return err
			}
			if len(args) == 0 {
				return handleRunErr(m.Up())
			}
			steps, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return handleRunErr(m.UpSteps(steps))
		},
	})
	migrationCmd.AddCommand(&cobra.Command{
		Use:   "down [steps]",
		Short: "Run migration down",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := migrate.NewMigrate(migrate.Options{})
			if err != nil {
				return err
			}
			if len(args) == 0 {
				return handleRunErr(m.Down())
			}
			steps, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return handleRunErr(m.DownSteps(steps))
		},
	})
	migrationCmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "View migration status",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := migrate.NewMigrate(migrate.Options{})
			if err != nil {
				return err
			}
			s, err := m.Status()
			if err != nil {
				return handleRunErr(err)
			}
			fmt.Printf("Current migration: %s\n", color.CyanString(s.CurrentVersionFile))
			if s.Dirty {
				color.Red("There are dirty migrations. Please fix this before running any migrations.")
				return nil
			}
			if s.Pending {
				color.Yellow("There are pending migrations. Please run:")
				fmt.Printf("\n  $ zepto db migration up\n\n")
			} else {
				color.Green("Migrations are up to date ✅")
			}
			return nil
		},
	})
	return &migrationCmd
}
