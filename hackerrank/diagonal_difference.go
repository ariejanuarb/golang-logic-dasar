package main

import "fmt"

/*
input :
3 (n of array)
11 2 4 (value 1,2
4 5 6
10 8 -12

process:
sum across the primary diagonal
sum across the secodnary diagonal
absolute difference |4-15|=15

output:
15

constraints:
-100 <= arr[i][j]<= 100
*/

func main() {
	// membuat array 2 dimensi
	array := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	differencial := diagonalDifference(array)
	fmt.Println(differencial)

	fmt.Println(array[0][0], array[0][1], array[0][2])
	fmt.Println(array[1][0], array[1][1], array[1][2])
	fmt.Println(array[2][0], array[2][1], array[2][2])

}

func diagonalDifference(arr [][]int32) int32 {
	var result int32
	var kanan int32 = 0
	var kiri int32 = 0

	// looping baris
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == j {
				kanan = kanan + arr[i][j]
			}
			if i+j == len(arr)-1 {
				kiri = kiri + arr[i][j]
			}
		}
	}
	// looping kolom
	if kanan > kiri {
		result = kanan - kiri
	} else {
		result = kiri - kanan
	}
	return result
}
