package CodeTest

func Aritmatika(arr []int) bool{
	var selisihSekarang int
	selisih := arr[1] - arr[0]
	for i := 0; i < len(arr) -1; i++ {
		selisihSekarang = arr[i+1] - arr[i]
		if selisihSekarang != selisih {
			return false
		}
	}
	return true
}