package ftls

import (
	"fmt"
	"os"
	"sort"
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
	f, err := os.Open(".")
	if err != nil {
		return err
	}
	dirContent, err := f.ReadDir(-1)
	if err != nil {
		return err
	}
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
	for i, dir := range dirContent {
		fmt.Printf("%s", dir.Name())
		if i < len(dirContent)-1 {
			fmt.Printf("  ")
		} else {
			fmt.Println()
		}
	}
	return nil
}
