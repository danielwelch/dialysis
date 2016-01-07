// fmover moves files with a certain extension (or extensions)
// from a source directory to a destination directory

package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type file struct {
	dirpath string
	name    string
}

func (f *file) fullPath() string {
	return filepath.Abs(dirpath + name)
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
		e, err := extension(f.Name())
		if err != nil {
			log.Printf("Invalid extension: %s", f.Name())
		}
		if e == ext {
			append(result, file{dirpath, f.Name()})
		}
	}
	return result
}

func validExtension(ext string) bool {
	if !strings.HasPrefix(ext, ".") {
		return false
	}
	return true
}

func extension(filename string) (string, error) {
	// get the file's extension
	i := strings.LastIndex(filename, ".")

	ext := filename[i:]
	if !validExtension(ext) {
		return "", errors.New("Invalid extension.")
	}
}

// move moves the given files to the new, given path.
// The new path will be the provided newdirpath joined
// with the original file name, as determined by
// os.FileInfo.Name()
func move(files []file, newdirpath string) error {
	for _, f := range files {
		if err := os.Rename(
			f.fullPath(),
			filepath.Abs(f.name+newdirpath),
		); err != nil {
			return err
		}
	}
	return nil
}

func getArgs() (string, string, string) {
	if len(os.Args) != 3 {
		log.Fatal("Error parsing command line arguments.\nRequires three args: SOURCE DESTINATION EXTENSION")
	}
	return os.Args[0], os.Args[1], os.Args[2]
}

func main() {

	// first arg = source dir
	// second arg = destination dir
	// third arg = file extension to match
	src, dst, ext := getArgs()

	start, err := filepath.Abs(src)
	if err != nil {
		log.Fatal(err)
	}
	end, err := filepath.Abs(dst)
	if err != nil {
		log.Fatal(err)
	}
	files, err := filterDir(start, ext)
	if err != nil {
		log.Fatal(err)
	}
	if err = move(files, dst); err != nil {
		log.Fatal(err)
	}
}
