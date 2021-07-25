package commands

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	zeptocli "github.com/go-zepto/zepto-cli"
	"github.com/go-zepto/zepto-cli/utils"
	"github.com/spf13/cobra"
)

var DEFAULT_TMPL_MODULE_PATH = "github.com/go-zepto/templates/default"
var DEFAULT_DOCKER_GO_ZEPTO_PATH = "github.com/go-zepto/zepto-cli/cmd/zepto"

var projectName string

func NpmInstall(dir string) {
	command := exec.Command("npm", "--silent", "--no-progress", "install")
	command.Dir = dir
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func GoModTidy(dir string) {
	command := exec.Command("go", "mod", "tidy")
	command.Dir = dir
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func ExecuteWeb(templates embed.FS, args []string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	fmt.Println("Creating web project...")
	s.Start()
	projectName := args[0]
	projectDir := "./" + projectName

	err := utils.CopyDirFromEmbed(templates, "_templates/web", projectDir)
	if err != nil {
		panic(err)
	}
	replaceFunc := func(c string) string {
		return strings.Replace(c, DEFAULT_TMPL_MODULE_PATH, projectName, -1)
	}
	err = filepath.Walk(projectDir, ReplaceWalk(projectDir, replaceFunc))
	if err != nil {
		panic(err)
	}
	utils.ReplaceTextOnFile(path.Join(projectDir, "Dockerfile"), DEFAULT_DOCKER_GO_ZEPTO_PATH, DEFAULT_DOCKER_GO_ZEPTO_PATH+"@v"+zeptocli.VERSION)
	fmt.Println("Preparing go mod...")
	s.Start()
	// Renaming files which can't be embed
	err = os.Rename(path.Join(projectDir, "go.mod_"), path.Join(projectDir, "go.mod"))
	if err != nil {
		panic(err)
	}
	err = os.Rename(path.Join(projectDir, "gitignore"), path.Join(projectDir, ".gitignore"))
	if err != nil {
		panic(err)
	}
	err = os.Rename(path.Join(projectDir, "dockerignore"), path.Join(projectDir, ".dockerignore"))
	if err != nil {
		panic(err)
	}
	// Run Go Mod Tidy
	GoModTidy(projectDir)
	s.Stop()
	fmt.Println("Installing npm packages...")
	s.Start()
	NpmInstall(projectDir)
	s.Stop()
	color.Green("Finished! Your project is ready.")
}

func NewCreateProjectCmd(templates embed.FS) *cobra.Command {
	cmd := cobra.Command{
		Use:   "new [name]",
		Short: "Create a new zepto project",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			utils.WarnVersion()
			ExecuteWeb(templates, args)
		},
	}
	return &cmd
}
