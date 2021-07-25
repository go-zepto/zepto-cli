package utils

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
)

func CopyDirFromEmbed(embedFs embed.FS, src string, dest string) error {
	f, err := embedFs.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}

	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	err = os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	files, err := embedFs.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {

		if f.IsDir() {

			err = CopyDirFromEmbed(embedFs, src+"/"+f.Name(), dest+"/"+f.Name())
			if err != nil {
				return err
			}

		}

		if !f.IsDir() {

			content, err := embedFs.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err

			}

			err = ioutil.WriteFile(dest+"/"+f.Name(), content, 0755)
			if err != nil {
				return err

			}

		}

	}

	return nil
}
