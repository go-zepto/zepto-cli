package zeptocli

import (
    "fmt"
    "github.com/markbates/pkger"
    "io"
    "os"
)

func PkgerCopyFile(source string, dest string) (err error) {
    sourcefile, err := pkger.Open(source)
    if err != nil {
        return err
    }

    defer sourcefile.Close()

    destfile, err := os.Create(dest)
    if err != nil {
        return err
    }

    defer destfile.Close()

    _, err = io.Copy(destfile, sourcefile)
    if err == nil {
        sourceinfo, err := pkger.Stat(source)
        if err != nil {
            err = os.Chmod(dest, sourceinfo.Mode())
        }

    }

    return
}


func PkgerCopyDir(source string, dest string) (err error) {

    // get properties of source dir
    sourceinfo, err := pkger.Stat(source)
    if err != nil {
        return err
    }

    // create dest dir

    err = os.MkdirAll(dest, sourceinfo.Mode())
    if err != nil {
        return err
    }

    directory, _ := pkger.Open(source)

    objects, err := directory.Readdir(-1)

    for _, obj := range objects {

        sourcefilepointer := source + "/" + obj.Name()
        destinationfilepointer := dest + "/" + obj.Name()

        if obj.IsDir() {
            // create sub-directories - recursively
            err = PkgerCopyDir(sourcefilepointer, destinationfilepointer)
            if err != nil {
                fmt.Println(err)
            }
        } else {
            // perform copy
            err = PkgerCopyFile(sourcefilepointer, destinationfilepointer)
            if err != nil {
                fmt.Println(err)
            }
        }

    }
    return
}
