package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	// цикл для поддержки множественного ввода
	for i := 1; i < len(arguments); i++ {
		s := arguments[i]
		ret := matchNameSur(s)
		fmt.Println(s, ": ", ret)
	}
}
