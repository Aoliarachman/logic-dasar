package array01


func NewStringArray(baris int, kolom int) [][]string {
	result := make([][]string, baris) //line berisi tentang membuat slice sepanjang "baris"
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


