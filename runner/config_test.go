package runner

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

const (
	bin = `./tmp/main`
	cmd = "go build -o ./tmp/main ."
)

func getWindowsConfig() Config {
	build := CfgBuild{
		Cmd:          "go build -o ./tmp/main .",
		Bin:          "./tmp/main",
		Log:          "build-errors.log",
		IncludeExt:   []string{"go", "tpl", "tmpl", "html"},
		ExcludeDir:   []string{"assets", "tmp", "vendor", "testdata"},
		ExcludeRegex: []string{"_test.go"},
		Delay:        1000,
		StopOnError:  true,
	}
	if runtime.GOOS == "windows" {
		build.Bin = bin
		build.Cmd = cmd
	}

	return Config{
		Root:        ".",
		TmpDir:      "tmp",
		TestDataDir: "testdata",
		Build:       build,
	}
}

func TestBinCmdPath(t *testing.T) {

	var err error

	c := getWindowsConfig()
	err = c.preprocess()
	if err != nil {
		t.Fatal(err)
	}

	if runtime.GOOS == "windows" {

		if !strings.HasSuffix(c.Build.Bin, "exe") {
			t.Fail()
		}

		if !strings.Contains(c.Build.Bin, "exe") {
			t.Fail()
		}
	} else {

		if strings.HasSuffix(c.Build.Bin, "exe") {
			t.Fail()
		}

		if strings.Contains(c.Build.Bin, "exe") {
			t.Fail()
		}
	}
}

func TestReadConfByName(t *testing.T) {
	_ = os.Unsetenv(airWd)
	Config, _ := readConfByName(dftTOML)
	if Config != nil {
		t.Fatalf("expect Config is nil,but get a not nil Config")
	}
}

func TestConfPreprocess(t *testing.T) {
	_ = os.Setenv(airWd, "_testdata/toml")
	df := defaultConfig()
	err := df.preprocess()
	if err != nil {
		t.Fatalf("preprocess error %v", err)
	}
	suffix := "/_testdata/toml/tmp/main"
	binPath := df.Build.Bin
	if !strings.HasSuffix(binPath, suffix) {
		t.Fatalf("bin path is %s, but not have suffix  %s.", binPath, suffix)
	}
}

func TestReadConfigWithWrongPath(t *testing.T) {
	c, err := readConfig("xxxx")
	if err == nil {
		t.Fatal("need throw a error")
	}
	if c != nil {
		t.Fatal("expect is nil but got a conf")
	}
}
