package commands

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-zepto/zepto-cli/runner"
	"github.com/go-zepto/zepto-cli/utils"
	"github.com/spf13/cobra"
)

func isZeptoProject() bool {
	b, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return false
	}
	goModContent := string(b)
	return strings.Contains(goModContent, "github.com/go-zepto/zepto")
}

func watch() {
	cfg := &runner.Config{
		Root:        ".",
		TmpDir:      "tmp",
		TestDataDir: "testdata",
		Build: runner.CfgBuild{
			Cmd:          "go build -o ./tmp/main *.go",
			Bin:          "tmp/main",
			IncludeExt:   []string{"go"},
			ExcludeDir:   []string{"tmp", "testdata", "node_modules", "build", "templates", "public"},
			ExcludeRegex: []string{"_test.go"},
			StopOnError:  true,
		},
	}
	if runtime.GOOS == runner.PlatformWindows {
		cfg.Build.Bin = "tmp\\main"
		if files, err := filepath.Glob("*.go"); err == nil {
			cfg.Build.Cmd = "go build -o ./tmp/main.exe " + strings.Join(files, " ")
		}
	}
	r, err := runner.NewEngineFromConfig(cfg, false)
	r.Run()
	if err != nil {
		panic(err)
	}
}

func ExecuteDev(args []string) {
	if !isZeptoProject() {
		log.Fatal("zepto dev failed: your current working dir is not a zepto project")
	}
	color.Green("Starting development server...")
	watch()
}

func NewDevCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "dev",
		Short: "Start a zepto project in development mode",
		Run: func(cmd *cobra.Command, args []string) {
			utils.WarnVersion()
			ExecuteDev(args)
		},
	}
	return &cmd
}
