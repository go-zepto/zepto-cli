package commands

import (
    "io/ioutil"
    "os"
    "path/filepath"
)

func ReplaceWalk(path string, f func(string) string) func(string, os.FileInfo, error) error {
	replaceFunc := func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !!fi.IsDir() {
			return nil //
		}
		matched, err := filepath.Match("*", fi.Name())
		if err != nil {
			panic(err)
			return err
		}
		if matched {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			newContents := f(string(read))

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}
		}
		return nil
	}
	return replaceFunc
}
