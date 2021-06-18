package CodeTest

import (
	strconv "strconv"
)

func PasanganTerbesar(angka int)string  {
	strangka := strconv.Itoa(angka)
	var terbesar = strangka[0:2]
	for i := 0; i <= len(strangka) -2 ; i++ {
		var cek = strangka[i:i+2]
		if cek > terbesar {
			terbesar = cek
		}
	}
	return terbesar
}
