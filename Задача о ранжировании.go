/*
Задача о ранжировании
Дана последовательность целых чисел с длиной N. Нужно найти m максимальных элементов (m << N — много меньше)

func topM(collection []int, m uint) []int

Что хочется увидеть: алгоритм со сложностью меньшей, чем O(N*log(N)) по времени. 
*/

package main

import (
	"fmt"
	"sort"
	)

func topM(collection []int, m uint) []int {
	sort.Ints(collection)
	startPos := uint(len(collection)) - m
	return collection[startPos :]
}

func main() {
	currentCollection := []int{ 10, 67, 78, 99, 68, 23, 3456, 754, 99, 234 }
	fmt.Println(currentCollection)
	fmt.Println(topM(currentCollection, 3))
}