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
	Dir                  string
	NoFlags              bool
	IgnoreEntries        bool
	Recursive            bool
	LongListingFormat    bool
	ReverseSort          bool
	ModificationTimeSort bool
}

func Help() {
	fmt.Println("Help")
}

func ProcessArgs(args ...string) {
	f := new(Flags)
	if len(args) == 1 {
		f.NoFlags = true
		f.IgnoreEntries = true
		f.Dir = "."
	}
	if err := PrintDirectory(f); err != nil {
		panic(err)
	}
}

func PrintDirectory(f *Flags) (err error) {
	var dirContent []fs.DirEntry
	if f.NoFlags {
		dirContent, err = getDirContent(f.Dir, alphabeticOrder)
	}
	if err != nil {
		return err
	}
	for i, dir := range dirContent {
		if f.IgnoreEntries && dir.Name()[0] == '.' {
			continue
		}
		if dir.Type().IsDir() {
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
	case 0:
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
