package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/go-zepto/zepto-cli/utils"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

func NpmBuild() {
	command := exec.Command("npm", "run", "build")
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

func GoBuild() {
	cmd := []string{"build", "-o", "build/app"}
	matches, err := filepath.Glob("*.go")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(matches); i++ {
		cmd = append(cmd, matches[i])
	}
	command := exec.Command("go", cmd...)
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		panic(err)
	}
}

func CopyResources() {
	err := copy.Copy("templates", "build/templates")
	if err != nil {
		panic(err)
	}
	err = copy.Copy("public", "build/public")
	if err != nil {
		panic(err)
	}
}

func ExecuteBuild(args []string) {
	if !isZeptoProject() {
		log.Fatal("zepto build failed: your current working dir is not a zepto project")
	}
	color.Cyan("Building zepto project ready for production...")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	os.Setenv("NODE_ENV", "production")
	os.Setenv("ZEPTO_ENV", "production")
	s.Start()
	fmt.Println("cleaning...")
	err := os.RemoveAll("public/build")
	if err != nil {
		panic(err)
	}
	err = os.RemoveAll("build")
	if err != nil {
		panic(err)
	}
	s.Stop()
	fmt.Println("npm build...")
	s.Start()
	NpmBuild()
	s.Stop()
	_ = os.Mkdir("build", os.ModePerm)
	fmt.Println("go build...")
	s.Start()
	GoBuild()
	s.Stop()
	fmt.Println("copying templates and public folder...")
	s.Start()
	CopyResources()
	s.Stop()
	color.Green("Success!")
	color.Green("Your zepto project is ready for production in build folder")
	fmt.Println("To start your project, just run:")
	fmt.Println("\ncd build && ./app")
}

func NewBuildCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "build",
		Short: "Build a zepto project ready for production",
		Run: func(cmd *cobra.Command, args []string) {
			utils.WarnVersion()
			ExecuteBuild(args)
		},
	}
	return &cmd
}
