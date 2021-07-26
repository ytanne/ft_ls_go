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
		PrintDirectory(f, ".")
		return
	}

	for _, dir := range directories {
		if err := PrintDirectory(f, dir); err != nil {
			panic(err)
		}
	}
}

func PrintDirectory(f *Flags, dir string) error {
	var directories []string
	dirContent, err := getDirContent(dir, alphabeticOrder)
	if err != nil {
		return err
	}

	if f.IncludeEntries {
		fmt.Printf(Purple("."))
		fmt.Print("  ")
		fmt.Printf(Purple(".."))
		fmt.Print("  ")
	}

	for i, dir := range dirContent {
		if !f.IncludeEntries && dir.Name()[0] == '.' {
			continue
		}
		if dir.Type().IsDir() {
			if f.Recursive {
				directories = append(directories, dir.Name())
			}
			fmt.Printf(Purple(dir.Name()))
		} else {
			fmt.Printf("%s", dir.Name())
		}
		if i < len(dirContent)-1 {
			fmt.Printf("  ")
		} else {
			fmt.Println()
		}
	}

	for _, intDir := range directories {
		PrintDirectory(f, intDir)
	}

	return nil
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
