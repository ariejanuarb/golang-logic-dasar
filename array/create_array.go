package array

func NewStringArray(baris int, kolom int) [][]string { // function NewString array, parameter baris int, kolom int, return value array 2 dimensi tipe string [][]string
	result := make([][]string, baris) // membuat slice, length baris
	for i := range result {
		result[i] = make([]string, kolom)
	}
	return result
}

func NewNumberArray(baris int, kolom int) [][]int32 {
	result := make([][]int32, baris)
	for i := range result {
		result[i] = make([]int32, kolom)
	}
	return result
}
