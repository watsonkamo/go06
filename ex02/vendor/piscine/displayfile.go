package piscine

import (
	"ft"
	"os"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func DisplayFile() {
	args := os.Args
	argCount := 0

	for range args {
		argCount++
	}
	if argCount != 2 {
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		return
	}
	defer file.Close()
	var buf [1024]byte
	for {
		n, err := file.Read(buf[:])
		os.Stdout.Write(buf[:n])
		if err != nil {
			break
		}
	}
}
