package main

// сортировка по любому полю среза структур, без обёртки или использования sort.Slice
import (
	"fmt"
	"os"
	"sort"
)

var SORTTYPE = ""

type Friend struct {
	Name    string
	Surname string
	Age     uint
}

type Friends []Friend

func (f Friends) Less(i, j int) bool {
	switch SORTTYPE {
	case "Name":
		return f[i].Name < f[j].Name
	case "Age":
		return f[i].Age < f[j].Age
	default:
		if f[i].Surname == f[j].Surname {
			return f[i].Name < f[j].Name
		}
		return f[i].Surname < f[j].Surname
	}
}

func (f Friends) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Friends) Len() int {
	return len(f)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("not enough arguments")
		return
	}
	SORTTYPE = os.Args[1]
	fs := Friends{Friend{"Bob", "Joy", 13}, Friend{"Lisa", "Brett", 19}, Friend{"Bob", "Ross", 56}, Friend{"Melisa", "Clay", 29}, Friend{"Milton", "Clay", 34}}
	fmt.Println(fs)
	sort.Sort(fs)
	for i, item := range fs {
		fmt.Println(i+1, item)
	}
}
