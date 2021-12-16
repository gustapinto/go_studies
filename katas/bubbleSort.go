package main

import (
	"fmt"
	"sort"
	"main/utils/tests"
)

func BubbleSort(vector [10]int) [10]int {
	for i := 0; i < len(vector); i++ {
		for j := 0; j < (len(vector) - 1); j++ {
			if vector[j] <= vector[j + 1] {
				continue
			}

			aux := vector[j]

			vector[j] = vector[j + 1]
			vector[j + 1] = aux
		}
	}

	return vector
}

func test(vector [10]int, sorted_vector [10]int) {
	expected_vector := vector
	expected_vector_slice := expected_vector[:]

	sort.Ints(expected_vector_slice)

	if expected_vector == sorted_vector {
		tests.PrintSuccess(sorted_vector)
	} else {
		tests.PrintFail(sorted_vector)
	}

	fmt.Println(string("\033[0m"))
}

func main() {
	vector1 := [10]int {1, 34, 2, 4, 5, 78, -3, 8, 10, -12}
	vector2 := [10]int {9, 8, 7, 6, 5, 3, 4, 2, 0, 1}
	vector3 := [10]int {-7, -1, -9, 0, -1, -3, -5, -6, -10, 1}

	test(vector1, BubbleSort(vector1))
	test(vector2, BubbleSort(vector2))
	test(vector3, BubbleSort(vector3))
}
