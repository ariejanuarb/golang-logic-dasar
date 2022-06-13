package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal05arr00(n int) {
	// 1. create new empty array
	array0500 := array2.NewNumberArray(n, n)

	//2. insert array value (n,n) using for loop
	value := 3
	nt := n / 2
	for rows := 0; rows < len(array0500); rows++ {
		for column := 0; column < len(array0500); column++ {
			array0500[rows][column] = int32(value)
		}
		// 2.1 condition the row
		if rows < nt {
			value += 3
		} else {
			value -= 3
		}
	}
	// 3. print array0500 filled with value
	array2.PrintNumberArray(array0500)
}
