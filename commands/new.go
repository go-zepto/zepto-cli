package commands

import (
	"fmt"
	zeptocli "github.com/go-zepto/zepto-cli"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var DEFAULT_TMPL_MODULE_PATH = "github.com/go-zepto/templates/default"


func NpmInstall(dir string) {
	fmt.Println("Installing NPM libraries...")
	command := exec.Command("npm", "--silent" +
		"", "install")
	command.Dir = dir
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Start()
}

func ExecuteWeb(args []string) {
	projectDir := "./" + path.Base(args[1])
	err := zeptocli.PkgerCopyDir("/_templates/web", projectDir)
	if err != nil {
		panic(err)
	}
	replaceFunc := func(c string) string {
		return strings.Replace(c, DEFAULT_TMPL_MODULE_PATH, args[1], -1)
	}

	err = filepath.Walk(projectDir, ReplaceWalk(projectDir, replaceFunc))
	if err != nil {
		panic(err)
	}
	NpmInstall(projectDir)
}

var NewCmd = &cobra.Command{
	Use:   "new [web|ms] [name]",
	Short: "Create a new zepto project",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "web" {
			ExecuteWeb(args)
		} else if args[0] == "ms" {
			fmt.Println("Not implemented (ms)")
		} else {
			fmt.Println("Invalid project type. You should select ms or web")
		}
	},
}
