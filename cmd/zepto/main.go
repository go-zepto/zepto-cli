package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/go-zepto/zepto-cli/commands"
	"github.com/go-zepto/zepto-cli/commands/db"
	"github.com/spf13/cobra"
)

//go:embed _templates/*
var templates embed.FS

var rootCmd = &cobra.Command{
	Use:   "zepto",
	Short: "Zepto is a lightweight web framework",
	Long:  "Zepto is a lightweight web  framework written in golang.\nComplete documentation is available at https://go-zepto.github.io/zepto",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(commands.NewCreateProjectCmd(templates))
	rootCmd.AddCommand(commands.NewDevCmd())
	rootCmd.AddCommand(commands.NewBuildCmd())
	rootCmd.AddCommand(db.NewDBCmd())
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
