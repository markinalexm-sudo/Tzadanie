package main

import "fmt"

func main() {
	// Исходные данные
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	// 1. Делаем уникальные списки
	// Использую свою функцию, чтобы не повторять код
	uniqA := removeDuplicates(a)
	uniqB := removeDuplicates(b)

	fmt.Printf("%v, %v\n", uniqA, uniqB)

	// 2. Ищем пересечения (среди уникальных)
	common := findIntersection(uniqA, uniqB)
	fmt.Printf("%v\n", common)

	// 3. Соединяем все вместе
	// append(a, b...) распаковывает второй слайс
	combined := append(a, b...)
	final := removeDuplicates(combined)

	fmt.Printf("%v\n", final)
}

// Функция удаляет повторы через карту
func removeDuplicates(nums []int) []int {
	m := make(map[int]bool)
	var result []int

	for _, n := range nums {
		if !m[n] {
			m[n] = true
			result = append(result, n)
		}
	}
	return result
}

// Функция ищет общие элементы
func findIntersection(a, b []int) []int {
	m := make(map[int]bool)
	// Записываем первый массив в карту
	for _, n := range a {
		m[n] = true
	}

	var res []int
	for _, n := range b {
		// Проверяем, есть ли число во втором массиве
		if m[n] {
			res = append(res, n)
		}
	}
	return res
}