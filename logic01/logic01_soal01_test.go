package logic01

import (
	"fmt"
	"math"
	"testing"
)

/*
Soal_01
nilai n	10
index	0	1	2	3	4	5	6	7	8	9
value	1	3	2	3	3	3	4	3	5	3

what? print given sequence of number (1 dimmensional array with n:= 10)
who? 2 starting number : a = 1 & b = 3
when? 2 pattern : a = 1+1, 2+1, 3+1 & b = 3, 3, 3, 3
where? conditoning even index number (i%2 == 0) inside the loop multiple of 2 (0,2,4,6,8) : (a += 1)
why? to create 1 dimensional array using 1 for loop + if method
how? using for loop method with 2 condition (if & else) to create 2 pattern in 1 dimensional array
*/

func TestLogic01Soal01(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i += 1 {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestLogic01Soal0101(t *testing.T) {
	n := 10
	a := 1
	b := 3
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(b, "\t")
		}
	}
}

func TestLogic01Soal0102(t *testing.T) {
	n := 10
	c := 1
	d := 3
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(c, "\t")
			c++
		} else {
			fmt.Print(d, "\t")
		}
	}
}

func TestLogic01Soal0103(t *testing.T) {
	n := 10
	e := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(e, "\t")
			e++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestLogic01Soal0104(t *testing.T) {
	n := 10
	f := 1
	g := 3
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(f, "\t")
			f++
		} else {
			fmt.Print(g, "\t")
		}
	}
}

func TestLogic01Soal0105(t *testing.T) {
	n := 10
	h := 1
	j := 3
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(h, "\t")
			h++
		} else {
			fmt.Print(j, "\t")
		}
	}
}

func TestLogic01Soal0106(t *testing.T) {
	n := 10
	l := 1
	m := 3
	for a := 0; a < n; a++ {
		if a%2 == 0 {
			fmt.Print(l, "\t")
			l++
		} else {
			fmt.Print(m, "\t")
		}
	}
}

func TestLogic01Soal0107(t *testing.T) {
	n := 10
	o := 1
	p := 3
	for z := 0; z < n; z++ {
		if z%2 == 0 {
			fmt.Print(o, "\t")
			o++
		} else {
			fmt.Print(p, "\t")
		}
	}
}

func TestLogic01Soal0108(t *testing.T) {
	n := 10
	q := 1
	r := 3
	for b := 0; b < n; b++ {
		if b%2 == 0 {
			fmt.Print(q, "\t")
			q++
		} else {
			fmt.Print(r, "\t")
		}
	}
}

func TestLogic01Soal0109(t *testing.T) {
	n := 10
	a := 1
	z := 3
	for c := 0; c < n; c++ {
		if c%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(z, "\t")
		}
	}
}

func TestLogic01Soal0110(t *testing.T) {
	n := 10
	i := 1
	j := 3
	for a := 0; a < n; a++ {
		if a%2 == 0 {
			fmt.Print(i, "\t")
			i++
		} else {
			fmt.Print(j, "\t")
		}
	}
}

func TestLogic01Soal0201(t *testing.T) {
	n := 10
	c := 3
	d := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(c, "\t")
		} else {
			fmt.Print(d, "\t")
			d++
		}

	}
}

func TestLogic01Soal03(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i += 1 {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(99, "\t")
		}
	}
}

func TestLogic01Soal04(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i += 1 {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(777, "\t")
		}
	}
}

func TestLogic01Soal05(t *testing.T) {
	n := 15
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz,buzz")
		} else if i%3 == 0 {
			fmt.Print("fizz", "\t")
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

func TestLogic01Soal07(t *testing.T) {
	n := 10
	a := 4
	b := 3
	for i := 1; i <= n; i += 1 {
		if i%2 == 0 {
			fmt.Print(b, "\t")
			b += 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func TestLogic01Soal08(t *testing.T) {
	n := 10
	a := 2
	b := 5
	for i := 1; i <= n; i += 1 {
		if i%2 == 0 {
			fmt.Print(b, "\t")
			b *= 5
		} else {
			fmt.Print(a, "\t")
			a *= 2
		}
	}
}

// 1^1, 2^1, 3^1,   1^2, 2^2, 3^2,   1^3, 2^3, 3^3,   1^4, 2^4, 3^4
func TestLogic01Soal09(t *testing.T) {
	n := 12
	a := 1 // bilangan yg akan dipangkatkan
	b := 1 // pangkat bilangan
	for i := 1; i <= n; i++ {
		fmt.Print(math.Pow(float64(a), float64(b)), "\t") // print angka pertama 1^1 = 1
		if i%3 == 0 {                                     // kondisi pada index ke-3, sebelum operasi deret selanjutnya dilakukan
			a = 1  // a reset kembali ke=1 di tiap index 3
			b += 1 // pangkat naik +1 tiap index 3 : 1,1,1 2,2,2 3,3,3 4,4,4
		} else {
			a += 1 // tambah angka+1 a : 2,3,, 1,2,3,, 1,2,3
		}
	}
}

// 1x3, 1x2, 1x1, 999,		2x3, 2x2, 2x1, 999, 	3x3, 3x2, 3x1, 999
func TestLogic01Soal10(t *testing.T) {
	n := 12
	a := 1 // bilangan yg akan dikalikan
	b := 4 // pengali bilangan
	for i := 1; i <= n; i++ {
		if i%4 == 0 { // kondisi tiap index ke-4, untuk merubah kondisi (a) ke tahap selanjutnya, dan (b) ke kondisi awal
			fmt.Print(999, "\t")
			a += 1 // a dinaikan +1 agar bilangan jadi 2
			b = 4  // b direset ke 4 untuk dikurangi lagi, agar pengali jadi 3,2,1
		} else {
			b -= 1
			fmt.Print(a*b, "\t")
		}
	}
}
