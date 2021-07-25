package runner

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"

	"github.com/pelletier/go-toml"
)

const (
	dftTOML = ".air.toml"
	dftConf = ".air.conf"
	airWd   = "air_wd"
)

type Config struct {
	Root        string   `toml:"root"`
	TmpDir      string   `toml:"tmp_dir"`
	TestDataDir string   `toml:"testdata_dir"`
	Build       CfgBuild `toml:"build"`
	Color       CfgColor `toml:"color"`
	Log         CfgLog   `toml:"log"`
	Misc        CfgMisc  `toml:"misc"`
}

type CfgBuild struct {
	Cmd              string        `toml:"cmd"`
	Bin              string        `toml:"bin"`
	FullBin          string        `toml:"full_bin"`
	Log              string        `toml:"log"`
	IncludeExt       []string      `toml:"include_ext"`
	ExcludeDir       []string      `toml:"exclude_dir"`
	IncludeDir       []string      `toml:"include_dir"`
	ExcludeFile      []string      `toml:"exclude_file"`
	ExcludeRegex     []string      `toml:"exclude_regex"`
	ExcludeUnchanged bool          `toml:"exclude_unchanged"`
	FollowSymlink    bool          `toml:"follow_symlink"`
	Delay            int           `toml:"delay"`
	StopOnError      bool          `toml:"stop_on_error"`
	SendInterrupt    bool          `toml:"send_interrupt"`
	KillDelay        time.Duration `toml:"kill_delay"`
	regexCompiled    []*regexp.Regexp
}

func (c *CfgBuild) RegexCompiled() ([]*regexp.Regexp, error) {
	if len(c.ExcludeRegex) > 0 && len(c.regexCompiled) == 0 {
		c.regexCompiled = make([]*regexp.Regexp, 0, len(c.ExcludeRegex))
		for _, s := range c.ExcludeRegex {
			re, err := regexp.Compile(s)
			if err != nil {
				return nil, err
			}
			c.regexCompiled = append(c.regexCompiled, re)
		}
	}
	return c.regexCompiled, nil
}

type CfgLog struct {
	AddTime bool `toml:"time"`
}

type CfgColor struct {
	Main    string `toml:"main"`
	Watcher string `toml:"watcher"`
	Build   string `toml:"build"`
	Runner  string `toml:"runner"`
	App     string `toml:"app"`
}

type CfgMisc struct {
	CleanOnExit bool `toml:"clean_on_exit"`
}

func readConfByName(name string) (*Config, error) {
	var path string
	if wd := os.Getenv(airWd); wd != "" {
		path = filepath.Join(wd, name)
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(wd, name)
	}
	cfg, err := readConfig(path)
	return cfg, err
}

func defaultConfig() Config {
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
	if runtime.GOOS == PlatformWindows {
		build.Bin = `tmp\main.exe`
		build.Cmd = "go build -o ./tmp/main.exe ."
	}
	log := CfgLog{
		AddTime: false,
	}
	color := CfgColor{
		Main:    "magenta",
		Watcher: "cyan",
		Build:   "yellow",
		Runner:  "green",
	}
	misc := CfgMisc{
		CleanOnExit: false,
	}
	return Config{
		Root:        ".",
		TmpDir:      "tmp",
		TestDataDir: "testdata",
		Build:       build,
		Color:       color,
		Log:         log,
		Misc:        misc,
	}
}

func readConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err = toml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func readConfigOrDefault(path string) (*Config, error) {
	dftCfg := defaultConfig()
	cfg, err := readConfig(path)
	if err != nil {
		return &dftCfg, err
	}

	return cfg, nil
}

func (c *Config) preprocess() error {
	var err error
	cwd := os.Getenv(airWd)
	if cwd != "" {
		if err = os.Chdir(cwd); err != nil {
			return err
		}
		c.Root = cwd
	}
	c.Root, err = expandPath(c.Root)
	if c.TmpDir == "" {
		c.TmpDir = "tmp"
	}
	if c.TestDataDir == "" {
		c.TestDataDir = "testdata"
	}
	if err != nil {
		return err
	}
	ed := c.Build.ExcludeDir
	for i := range ed {
		ed[i] = cleanPath(ed[i])
	}

	adaptToVariousPlatforms(c)

	c.Build.ExcludeDir = ed
	if len(c.Build.FullBin) > 0 {
		c.Build.Bin = c.Build.FullBin
		return err
	}
	// Fix windows CMD processor
	// CMD will not recognize relative path like ./tmp/server
	c.Build.Bin, err = filepath.Abs(c.Build.Bin)
	return err
}

func (c *Config) colorInfo() map[string]string {
	return map[string]string{
		"main":    c.Color.Main,
		"build":   c.Color.Build,
		"runner":  c.Color.Runner,
		"watcher": c.Color.Watcher,
	}
}

func (c *Config) buildLogPath() string {
	return filepath.Join(c.tmpPath(), c.Build.Log)
}

func (c *Config) buildDelay() time.Duration {
	return time.Duration(c.Build.Delay) * time.Millisecond
}

func (c *Config) binPath() string {
	return filepath.Join(c.Root, c.Build.Bin)
}

func (c *Config) tmpPath() string {
	return filepath.Join(c.Root, c.TmpDir)
}

func (c *Config) TestDataPath() string {
	return filepath.Join(c.Root, c.TestDataDir)
}

func (c *Config) rel(path string) string {
	s, err := filepath.Rel(c.Root, path)
	if err != nil {
		return ""
	}
	return s
}
