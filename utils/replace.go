package utils

import (
	"io/ioutil"
	"strings"
)

func ReplaceTextOnFile(path, oldText, newText string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), oldText, newText, -1)
	return ioutil.WriteFile(path, []byte(newContents), 0)
}
