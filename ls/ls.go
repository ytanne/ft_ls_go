package ftls

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
)

const (
	alphabeticOrder = iota
)

func Help() {
	fmt.Println("Help")
}

func ProcessArgs(args ...string) {
	if len(args) == 1 {
		if err := PrintDirectory(); err != nil {
			panic(err)
		}
		return
	}
}

func PrintDirectory() error {
	dirContent, err := getDirContent(".", alphabeticOrder)
	if err != nil {
		return err
	}
	for i, dir := range dirContent {
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
