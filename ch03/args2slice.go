package main

import (
	"fmt"
	"os"
)

type Args struct {
	Index int
	Word  string
}

// анписать утилиту, которая преобразует аргументы командной строки в срез структур из индекса и содержания позиции
// строка для теста
// go run args2slice.go one two three four five six brown --==++ =^-^= ten

func main() {
	arguments := os.Args
	args := []Args{}
	for i, item := range arguments {
		args = append(args, Args{Index: i, Word: item})
	}
	fmt.Println(args)
}
