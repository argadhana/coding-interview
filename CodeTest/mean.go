package CodeTest

func Mean(arr []int)int  {
	result := 0
	for _, v := range arr {
		result += v
	}
	 result = result / len(arr)
	 return result
}