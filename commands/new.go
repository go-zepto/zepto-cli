package commands

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	zeptocli "github.com/go-zepto/zepto-cli"
	"github.com/go-zepto/zepto-cli/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var DEFAULT_TMPL_MODULE_PATH = "github.com/go-zepto/templates/default"


func NpmInstall(dir string) {
	command := exec.Command("npm", "--silent", "--no-progress", "install")
	command.Dir = dir
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func ExecuteWeb(args []string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	fmt.Println("Creating web project...")
	projectDir := "./" + path.Base(args[1])
	err := zeptocli.PkgerCopyDir("/_templates/web", projectDir)
	if err != nil {
		panic(err)
	}
	replaceFunc := func(c string) string {
		return strings.Replace(c, DEFAULT_TMPL_MODULE_PATH, args[1], -1)
	}
	s.Stop()
	err = filepath.Walk(projectDir, ReplaceWalk(projectDir, replaceFunc))
	if err != nil {
		panic(err)
	}
	fmt.Println("Installing npm packages...")
	s.Start()
	NpmInstall(projectDir)
	s.Stop()
	color.Green("Finished! Your project is ready.");
}

var NewCmd = &cobra.Command{
	Use:   "new [web|ms] [name]",
	Short: "Create a new zepto project",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		utils.WarnVersion()
		if args[0] == "web" {
			ExecuteWeb(args)
		} else if args[0] == "ms" {
			fmt.Println("Not implemented (ms)")
		} else {
			fmt.Println("Invalid project type. You should select ms or web")
		}
	},
}
