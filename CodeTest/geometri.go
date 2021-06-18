package CodeTest

func Geometri(arr []int) bool{
	rasio := arr[1] / arr[0]
	var rasioSekarang int
	for i := 0; i < len(arr) -1; i++ {
		rasioSekarang = arr[i+1] / arr[i]
		if rasioSekarang != rasio {
			return false
		}
	}
	return true
}
