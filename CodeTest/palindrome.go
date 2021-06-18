package CodeTest

func PalindromeReverse(kata string) bool {
	var temp string
	for i := len(kata) - 1; i >=0; i-- {
		temp = temp + string(kata[i])
	}
	return temp == kata
}

func PalindromeWoTemp(value string) bool{
	for i:= 0; i < len(value) / 2 ; i++  {
		indexAwal := i
		indexAkhir := len(value) - i -1

		if string(value[indexAwal]) != string(value[indexAkhir]){
			return false
		}
	}
	return true;
}

func PalindromeRecursive(isi string, i int) bool{
	if i < len(isi) /2 {
		if string(isi[i]) != string(isi[len(isi) - i -1]){
			return false
		} else {
			return PalindromeRecursive(isi, i+ 1)
		}
	}else{
		return true
	}
}