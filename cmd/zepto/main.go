package main

import (
	"fmt"
	"github.com/go-zepto/zepto-cli/commands"
	"github.com/spf13/cobra"
	"os"
)


var rootCmd = &cobra.Command{
	Use:   "zepto",
	Short: "Zepto is a lightweight web & microservices framework",
	Long: "Zepto is a lightweight framework for the development of microservices & web services in golang.\nComplete documentation is available at https://go-zepto.github.io/zepto",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func main() {
	rootCmd.AddCommand(commands.NewCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}