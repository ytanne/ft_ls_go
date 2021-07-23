package main

import (
	"os"

	ftls "github.com/ytanne/ft_ls_go/ls"
)

func main() {
	ftls.ProcessArgs(os.Args...)
}
