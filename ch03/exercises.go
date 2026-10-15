package main

import (
	"fmt"
)

// написать функцию, преобращующую массив в карту
func array2Map(arrptr *[]string) map[int]string {
	mp := make(map[int]string)
	for i, item := range *arrptr {
		mp[i] = item
	}
	return mp
}

// написать функцию, преобразующую карту в два среза (ключей, значений с общим идексом)
func map2Arrays(mp map[int]string) ([]int, []string) {
	keys := make([]int, 0, len(mp))
	values := make([]string, 0, len(mp))
	for k, v := range mp {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func main() {
	fmt.Println("Первое задание")
	var arr = []string{"zero", "one", "two", "three", "four", "five"}
	mp := array2Map(&arr)
	fmt.Println(mp)

	fmt.Println("Второе задание")
	k, v := map2Arrays(mp)
	for i := 0; i < len(k); i++ {
		fmt.Printf("index: %d; key: %d, value: %s\n", i, k[i], v[i])
	}
}
