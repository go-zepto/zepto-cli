package utils

import (
    "fmt"
    "github.com/fatih/color"
    zeptocli "github.com/go-zepto/zepto-cli"
    "github.com/tcnksm/go-latest"
)

func WarnVersion() {
    fmt.Println("zepto-cli@" + zeptocli.VERSION)
    githubTag := &latest.GithubTag{
        Owner: "go-zepto",
        Repository: "zepto-cli",
    }

    res, _ := latest.Check(githubTag, zeptocli.VERSION)
    if res != nil && res.Outdated {
        color.Yellow("A new zepto-cli version (%s) is available.\n", res.Current)
        fmt.Printf("Please, consider upgrade:\n\tgo get -u github.com/go-zepto/zepto-cli/cmd/zepto\n\n")
    }
}