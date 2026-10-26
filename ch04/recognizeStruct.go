package main

// создание функции узнающей структуру
import (
	"fmt"
)

type Friend struct {
	Name    string
	Surname string
	Age     uint
}

type Enemy struct {
	StupidName string
	Age        uint
	Life       bool
}

type Human struct {
	Name string
}

func recognizer(s interface{}) {
	switch v := s.(type) {
	case Friend:
		fmt.Println("It's Friend:", v)
	case Enemy:
		fmt.Println("It's Enemy:", v)
	default:
		fmt.Println("We don’t know each other yet.")
	}
}

func main() {
	f := Friend{"Bob", "Ross", 54}
	e := Enemy{"Ford", 26, true}
	h := Human{"Oleg"}
	recognizer(f)
	recognizer(e)
	recognizer(h)
}
