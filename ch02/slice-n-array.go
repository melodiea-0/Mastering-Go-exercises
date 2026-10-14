package main

import (
	"fmt"
)

// объединяет два массива в новый срез
func unionArray2Slice(arr1 *[5]int, arr2 *[5]int) []int {
	newSliceFromArray := append(arr1[:], arr2[:]...)
	return newSliceFromArray
}

// объединяет два массива в новый массив
func unionArray2Array(arr1 [5]int, arr2 [5]int) [10]int {
	var newArray [10]int
	copy(newArray[:5], arr1[:])
	copy(newArray[5:], arr2[:])
	return newArray
}

// объединяет два среза в новый массив
func unionSlice2Array(sl1 []int, sl2 []int) ([10]int, error) {
	var newSlice [10]int
	if len(sl1)+len(sl2) != len(newSlice) {
		return newSlice, fmt.Errorf("ожидаем %d элементов, получено %d", len(newSlice), len(sl1)+len(sl2))
	}
	n := copy(newSlice[:], sl1)
	copy(newSlice[n:], sl2)
	return [10]int(newSlice), nil
}

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{6, 7, 8, 9, 10}
	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := []int{6, 7, 8, 9, 10}
	fmt.Println("Выводи исходные объекты")
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("sl1:", sl1)
	fmt.Println("sl2:", sl2)

	newSliceFromArray := unionArray2Slice(&arr1, &arr2)
	newSliceFromArray[0] *= 12
	newSliceFromArray[9] *= 12
	fmt.Println()
	fmt.Printf("arr1 + arr2 = %T: ", newSliceFromArray)
	fmt.Println(newSliceFromArray)

	newArrayFromArray := unionArray2Array(arr1, arr2)
	newArrayFromArray[0] *= 13
	newArrayFromArray[9] *= 13
	fmt.Println()
	fmt.Printf("arr1 + arr2 = %T: ", newArrayFromArray)
	fmt.Println(newArrayFromArray)

	newArrayFromSlice, _ := unionSlice2Array(sl1, sl2)
	newArrayFromSlice[0] *= 14
	newArrayFromSlice[9] *= 14
	fmt.Println()
	fmt.Printf("sl1 + sl2 = %T: ", newArrayFromSlice)
	fmt.Println(newArrayFromSlice)

	fmt.Println()
	fmt.Println("Убеждаемся, что исходные данные не изменились и мы создали новые объекты")
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("sl1:", sl1)
	fmt.Println("sl2:", sl2)
}
