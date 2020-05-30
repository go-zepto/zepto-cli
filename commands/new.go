package commands

import (
	"fmt"
	"github.com/fatih/color"
	zeptocli "github.com/go-zepto/zepto-cli"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"github.com/briandowns/spinner"
	"time"
	"github.com/tcnksm/go-latest"
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

		githubTag := &latest.GithubTag{
			Owner: "go-zepto",
			Repository: "zepto-cli",
		}

		res, _ := latest.Check(githubTag, zeptocli.VERSION)
		if res != nil && res.Outdated {
			fmt.Printf("%s is not latest. Please, consider upgrade to %s:\n go get -u github.com/go-zepto/zepto-cli/cmd/zepto", zeptocli.VERSION, res.Current)
		}



		if args[0] == "web" {
			ExecuteWeb(args)
		} else if args[0] == "ms" {
			fmt.Println("Not implemented (ms)")
		} else {
			fmt.Println("Invalid project type. You should select ms or web")
		}
	},
}
