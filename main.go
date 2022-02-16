package main

import "fmt"

//Функция "Чудные вхождения в массив"
func Solution(A []int) int {
	switch len(A) {
	case 0:
		return 0
	case 1:
		return A[0]
	default:
		var m map[int]bool = make(map[int]bool)
		for _, val := range A {
			m[val] = !m[val]
		}
		for val, res := range m {
			if res {
				return val
			}
		}
		return 0
	}

}

func main() {
	A := []int{9, 3, 9, 3, 9, 5, 9}
	fmt.Println(Solution(A))
}
