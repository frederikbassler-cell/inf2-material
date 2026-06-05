package arrays

import "fmt"

func Example_matrix() {
	m := [][]int{
		{10, 20, 30},
		{40, 50, 60},
		{70, 80, 90},
	}

	fmt.Println(m[0])
	fmt.Println(m[1])
	fmt.Println(m[2])
	fmt.Println()

	fmt.Println(m[0][0])
	fmt.Println(m[1][1])
	fmt.Println(m[2][2])
	fmt.Println()

	fmt.Println(m[0][2])
	fmt.Println(m[2][0])

	// Output:
	// [10 20 30]
	// [40 50 60]
	// [70 80 90]
	//
	// 10
	// 50
	// 90
	//
	// 30
	// 70
}
