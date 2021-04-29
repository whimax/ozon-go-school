/*
Задача о поиске максимума и минимума
Дана последовательность из целых чисел. Нужно найти минимальный и максимальный элемент в коллекции.

func minmax(collection []int) (int, int)

Что должно получиться:
1. Алгоритм за один проход с константной сложностью по памяти.
2. Разбиение на пары для оптимального количества сравнений.
3. Проверка того, что коллекция может быть пустой.
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func duration(msg string, start time.Time) {
	dur := time.Since(start)
	fmt.Printf("%v: %v\n", msg, dur.Nanoseconds())
}

func collectionGen(size int) []int {
	coll := make([]int, size, size)
	for i := 0; i < size; i++ {
		coll[i] = rand.Intn(size)
	}
	return coll
}

func minmaxStd(collection []int) (int, int) {
	startTime := time.Now()
	defer duration("minmax std sort", startTime)

	if len(collection) != 0 {
		sort.Ints(collection)
		return collection[0], collection[len(collection)-1]
	}
	return 0, 0
}

func minmax(collection []int) (int, int) {
	startTime := time.Now()
	defer duration("minmax", startTime)

	if len(collection) != 0 {

		max := collection[0]
		min := max
		for i := 1; i < len(collection); i++ {
			if collection[i] > max {
				max = collection[i]
			}
			if collection[i] < min {
				min = collection[i]
			}
		}
		return min, max
	}
	return 0, 0
}

func main() {
	curCollection := collectionGen(10000)
	fmt.Println(curCollection)
	fmt.Println(minmax(curCollection))
	fmt.Println(minmaxStd(curCollection))
}
