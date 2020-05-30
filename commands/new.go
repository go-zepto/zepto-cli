package commands

import (
	"fmt"
	"github.com/go-zepto/zepto-cli/utils"
	"github.com/spf13/cobra"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var DEFAULT_TMPL_MODULE_PATH = "github.com/go-zepto/templates/default"

func getTemplatePath() string {
	_, f, _, _ := runtime.Caller(0)
	fp := path.Join(path.Dir(f), "./../_templates/web")
	return fp
}

func ExecuteWeb(args []string) {
	filename := path.Base(args[1])
	err := utils.PkgerCopyDir(getTemplatePath(), "./" + filename)
	if err != nil {
		panic(err)
	}
	replaceFunc := func(c string) string {
		return strings.Replace(c, DEFAULT_TMPL_MODULE_PATH, args[1], -1)
	}

	err = filepath.Walk("./" + filename, ReplaceWalk("./" + filename, replaceFunc))
	if err != nil {
		panic(err)
	}
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
