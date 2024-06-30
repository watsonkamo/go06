package main

import (
	"os"
	"piscine"
)

const Notfound = "File name missing\n"
const Toomany = "Too many arguments\n"

func main() {
	args := os.Args
	if len(args) == 1 {
		piscine.PrintStr(Notfound)
	}
	if len(args) > 2 {
		piscine.PrintStr(Toomany)
	}
	piscine.DisplayFile()
}
