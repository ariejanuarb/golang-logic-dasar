package logic_sabtu

import (
	"fmt"
	"math"
	"testing"
)

func TestLogic01Soal01(t *testing.T) {
	n := 10
	angka := 1
	for i := 1; i < n; i += 1 {
		if i%2 == 0 {
			fmt.Print(i, "\t")
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

func TestSoal02(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i += 1 {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestSoal03(t *testing.T) {
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

func TestSoal04(t *testing.T) {
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

func TestSoal05(t *testing.T) {
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

func TestSoal06(t *testing.T) {
	n := 15
	nt := n / 2
	angka := 3

	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			/*x := math.Pow(float64(i), 2)
			fmt.Print(x, "\t")*/
			fmt.Print(i*i, "\t")
		} else {
			fmt.Print(angka, "\t")
		}
		if i <= nt {
			angka += 3
		} else {
			angka -= 3
		}
	}
}

func TestSoal07(t *testing.T) {
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

func TestSoal08(t *testing.T) {
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
func TestSoal09(t *testing.T) {
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
func TestSoal10(t *testing.T) {
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
