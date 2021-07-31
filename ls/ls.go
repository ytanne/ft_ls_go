package ftls

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
)

const (
	alphabeticOrder = iota
	reverseOrder
	modificationTimeOrder
)

type Flags struct {
	IncludeEntries       bool
	Recursive            bool
	LongListingFormat    bool
	ReverseSort          bool
	ModificationTimeSort bool
}

func Help() {
	fmt.Println("Help")
}

func ProcessArgs(args ...string) {
	f := &Flags{}
	var directories []string

	for _, arg := range args[1:] {
		if arg[0] == '-' {
			if len(arg) == 1 {
				panic("Cannot open directory '-'")
			}
			for _, letter := range arg[1:] {
				switch letter {
				case 'l':
					f.LongListingFormat = true
				case 'a':
					f.IncludeEntries = true
				case 'R':
					f.Recursive = true
				case 'r':
					f.ReverseSort = true
				case 't':
					f.ModificationTimeSort = true
				default:
					fmt.Printf("flag -%v is not implemented yet", letter)
				}
			}
		} else {
			directories = append(directories, arg)
		}
	}

	if len(directories) == 0 {
		objects := processDirArgs(f, ".")
		printObjects(objects)
		return
	}
	allObjects := make([]Object, 0)
	for _, directory := range directories {
		allObjects = append(allObjects, processDirArgs(f, directory)...)
	}
	printObjects(allObjects)
}

func getDirContent(dir string, order int) ([]fs.DirEntry, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	dirContent, err := f.ReadDir(-1)
	if err != nil {
		return nil, err
	}
	switch order {
	case alphabeticOrder:
		sort.Slice(dirContent, func(i, j int) bool {
			var a, b = dirContent[i].Name()[0], dirContent[j].Name()[0]
			if a >= 'A' && a <= 'Z' {
				a += ' '
			}
			if b >= 'A' && b <= 'Z' {
				b += ' '
			}
			return a < b
		})
	}
	return dirContent, nil
}

func processDirArgs(f *Flags, dir string) []Object {
	fd, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	fileEntries, err := fd.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	allObjects := make([]Object, 0, len(fileEntries))
	if f.IncludeEntries {
		allObjects = append(allObjects, getDotsInfo()...)
	}

	for _, entry := range fileEntries {
		if !f.IncludeEntries && entry.Name()[0] == '.' {
			continue
		}

	}
	return nil
}
