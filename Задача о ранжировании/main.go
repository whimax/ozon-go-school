/*
������ � ������������
���� ������������������ ����� ����� � ������ N. ����� ����� m ������������ ��������� (m << N � ����� ������)

func topM(collection []int, m uint) []int

��� ������� �������: �������� �� ���������� �������, ��� O(N*log(N)) �� �������. 
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