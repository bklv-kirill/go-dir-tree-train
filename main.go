package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	Name   string
	Size   int64
	Files  []*File
	IsFile bool
}

func dirTree(out io.Writer, path string, withFiles bool) (err error) {
	var dirs []*File = make([]*File, 0)

	var cls func(path string, pDir *File) error
	cls = func(path string, pDir *File) error {
		dFiles, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, dFile := range dFiles {
			dFileI, err := dFile.Info()
			if err != nil {
				return err
			}

			if dFile.IsDir() {
				var dir *File = &File{
					Name: dFileI.Name(),
				}

				if pDir != nil {
					pDir.Files = append(pDir.Files, dir)
				} else {
					dirs = append(dirs, dir)
				}

				err = cls(filepath.Join(path, dFileI.Name()), dir)
				if err != nil {
					return err
				}
			} else if withFiles {
				var file *File = &File{
					Name:   dFileI.Name(),
					Size:   dFileI.Size(),
					IsFile: true,
				}

				if pDir != nil {
					pDir.Files = append(pDir.Files, file)
				} else {
					dirs = append(dirs, file)
				}
			}
		}

		return nil
	}

	err = cls(path, nil)
	if err != nil {
		return fmt.Errorf("dataTree | closure: %v", err)
	}

	var possForSkip []int = make([]int, len(dirs))
	for i, dir := range dirs {
		var isLast bool = i == len(dirs)-1
		err = printDirTree(out, dir, 1, possForSkip, isLast)
		if err != nil {
			return fmt.Errorf("dataTree | print:  %v", err)
		}
	}

	return nil
}

func printDirTree(out io.Writer, file *File, pos int, possForSkip []int, isLast bool) error {
	var possForSkipCP []int = make([]int, len(possForSkip))
	copy(possForSkipCP, possForSkip)

	var pref string
	for i := 1; i < pos+1; i++ {
		if i == pos {
			if isLast {
				pref += "└───"
			} else {
				pref += "├───"
			}

			continue
		}

		if needSkip(i, possForSkipCP) {
			pref += "\t"
		} else {
			pref += "│\t"
		}
	}

	if isLast {
		possForSkipCP = append(possForSkipCP, pos)
	}

	switch file.IsFile {
	case true:
		if file.Size == 0 {
			_, err := out.Write([]byte(fmt.Sprintf("%s (%s)\n", pref+file.Name, "empty")))
			if err != nil {
				return fmt.Errorf("printDirTree: %v", err)
			}
		} else {
			_, err := out.Write([]byte(fmt.Sprintf("%s (%db)\n", pref+file.Name, file.Size)))
			if err != nil {
				return fmt.Errorf("printDirTree: %v", err)
			}
		}
	default:
		_, err := out.Write([]byte(fmt.Sprintf("%s\n", pref+file.Name)))
		if err != nil {
			return fmt.Errorf("printDirTree: %v", err)
		}
	}

	for index, childFile := range file.Files {
		isLast = index == len(file.Files)-1
		err := printDirTree(out, childFile, pos+1, possForSkipCP, isLast)
		if err != nil {
			return err
		}
	}

	return nil
}

func needSkip(pos int, possForSkipCP []int) (skip bool) {
	for _, posForSkipCP := range possForSkipCP {
		if pos == posForSkipCP {
			skip = true
		}
	}
	return
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	withFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, withFiles)
	if err != nil {
		panic(err)
	}
}
