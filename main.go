package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err)
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	var printTree func(path, init string, printFiles bool) error
	printTree = func(path, init string, printFiles bool) error {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		})

		if !printFiles {
			directories := make([]os.DirEntry, 0, len(entries))
			for _, entry := range entries {
				if entry.IsDir() {
					directories = append(directories, entry)
				}
			}
			entries = directories
		}

		for i, entry := range entries {
			eInfo, err := entry.Info()
			if err != nil {
				return err
			}

			var pref string
			if i+1 == len(entries) {
				pref = "└───"
			} else {
				pref = "├───"
			}

			if entry.IsDir() {
				fmt.Fprintf(out, "%s%s\n", init+pref, eInfo.Name())

				var pref2 string
				if i+1 == len(entries) {
					pref2 += init + "\t"
				} else {
					pref2 += init + "│\t"
				}

				err = printTree(filepath.Join(path, eInfo.Name()), pref2, printFiles)
				if err != nil {
					return err
				}
			} else if printFiles {
				var size string
				if eInfo.Size() != 0 {
					size = fmt.Sprintf("%db", eInfo.Size())
				} else {
					size = "empty"
				}
				fmt.Fprintf(out, "%s%s (%s)\n", init+pref, eInfo.Name(), size)
			}
		}

		return nil
	}

	err := printTree(path, "", printFiles)
	if err != nil {
		return err
	}

	return nil
}
