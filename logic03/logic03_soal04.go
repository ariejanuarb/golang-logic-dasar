package logic03

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic03Soal04arr00Step01(n int) { // create first 4 line that are dynamic
	//1 create new array
	array0400 := array2.NewNumberArray(n, n)
	//2.1 initialize variable and value : start with 3, +3 in new row. and -3 after middle constraint (n/2)
	number := 3
	middle := n / 2
	//2 create for-loop and use its variable as array index
	for rows := 0; rows < len(array0400); rows++ {
		for column := 0; column < len(array0400); column++ {
			//2.2 insert value to index (val. need parsing)
			if rows%4 == 0 { // rows 0,4,8 so on : start from 0 +4
				array0400[rows][column] = int32(number)
			} else if rows%4 == 1 && column == n-1 { // rows 1,5,9 so on & column n-1 : start from 1 +4
				array0400[rows][column] = int32(number)
			} else if rows%4 == 2 { // rows 2,6,10 so on : start from 2, +4 +4
				array0400[rows][column] = int32(number)
			} else if rows%4 == 3 && column == n-n { // rows 3,7,11 so on & column 0 : start from 3 +4
				array0400[rows][column] = int32(number)
			}
		}
		//2.3 write condition on rows to create number pattern
		if rows < middle {
			number += 3
		} else {
			number -= 3
		}
	}
	//3 print array
	array2.PrintNumberArrayZeroEmpty(array0400)
}
