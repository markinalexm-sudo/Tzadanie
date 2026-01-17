package main

import "fmt"

func main() {
	// Входные данные из условия
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	// 1. Убираем дубликаты
	cleanA := removeDuplicates(a)
	cleanB := removeDuplicates(b)

	fmt.Printf("%v, %v\n", cleanA, cleanB)

	// 2. Находим пересечения (ищем общие числа)
	common := findIntersection(cleanA, cleanB)
	fmt.Printf("%v\n", common)

	// 3. Объединяем оба списка в один уникальный
	combined := append(a, b...)
	uniqueCombined := removeDuplicates(combined)

	fmt.Printf("%v\n", uniqueCombined)
}

// removeDuplicates возвращает слайс без повторов
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, num := range nums {
		// Если числа еще не было в map, добавляем его
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

// findIntersection ищет общие элементы в двух списках
func findIntersection(a, b []int) []int {
	// Делаем "справочник" из первого списка
	checkMap := make(map[int]bool)
	for _, num := range a {
		checkMap[num] = true
	}

	var intersection []int
	for _, num := range b {
		// Если число из второго списка есть в справочнике
		if checkMap[num] {
			intersection = append(intersection, num)
		}
	}
	return intersection
}