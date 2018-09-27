package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// c tyle of string => int
type c map[string]int

func main() {
	// arguments are list of strings
	// default: .
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	// iterate over all directories
	for _, arg := range args {
		// creating an empty map
		count := make(c)
		data, err := tree(arg, "", count)
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
		fmt.Printf("\n%d directories, %d files\n", data["dirs"], data["files"])
	}
}

func tree(root, indent string, count c) (c, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return make(c), fmt.Errorf("could not stat %s: %v", root, err)
	}

	// we need to print it here to get files and dirs
	fmt.Println(fi.Name())

	// if it's file then return nil
	if !fi.IsDir() {
		return make(c), nil
	}

	// we want to read only dirs here
	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return make(c), fmt.Errorf("could not read dir %s: %v", root, err)
	}

	var names []string
	for _, fi := range fis {
		// let's count how many files and dirs we have in root
		if fi.IsDir() {
			count["dirs"]++
		} else {
			count["files"]++
		}
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
			continue
		}
	}

	// iterate over dirs to get files
	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			fmt.Printf(indent + "└──")
			add = "   "
		} else {
			fmt.Printf(indent + "├──")
		}

		_, err := tree(filepath.Join(root, name), indent+add, count)
		if err != nil {
			return make(c), err
		}
	}
	return count, nil
}
