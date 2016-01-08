// fmover moves files with a certain extension (or extensions)
// from a source directory to a destination directory

package main

import (
	// "errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	// "strings"
)

type file struct {
	dirpath string
	name    string
}

func (f *file) fullPath() string {
	return f.dirpath + f.name
}

//filterDirTwo returns file structs for just the files in a
// given directory that matches extension ext
func filterDir(dirpath, ext string) ([]file, error) {
	var result []file
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			log.Printf("Skipping %s: is a directory", f.Name())
		}
		e := filepath.Ext(f.Name())
		if err != nil {
			log.Println(err)
		}
		if e == ext {
			result = append(result, file{dirpath, f.Name()})
		}
	}
	return result, nil
}

// func validExtension(ext string) bool {
// 	if !strings.HasPrefix(ext, ".") {
// 		return false
// 	}
// 	return true
// }

// func extension(filename string) (string, error) {
// 	// get the file's extension
// 	i := strings.LastIndex(filename, ".")

// 	ext := filename[i:]
// 	if !validExtension(ext) {
// 		return "", errors.New("ERROR: Invalid extension: %s")
// 	}
// 	return ext, nil
// }

// move moves the given files to the new, given path.
// The new path will be the provided newdirpath joined
// with the original file name, as determined by
// os.FileInfo.Name()
func move(files []file, newdirpath string) error {
	for _, f := range files {
		path := newdirpath + f.name
		if err := os.Rename(
			f.fullPath(),
			path,
		); err != nil {
			return err
		}
	}
	return nil
}

func getArgs() (string, string, string) {
	if len(os.Args) != 4 {
		log.Fatal("Error parsing command line arguments.\nRequires three args: SOURCE DESTINATION EXTENSION")
	}
	return os.Args[1], os.Args[2], os.Args[3]
}

func main() {

	// first arg = source dir
	// second arg = destination dir
	// third arg = file extension to match
	src, dst, ext := getArgs()

	files, err := filterDir(src, ext)
	if err != nil {
		log.Fatal(err)
	}
	if err = move(files, dst); err != nil {
		log.Fatal(err)
	}
}
