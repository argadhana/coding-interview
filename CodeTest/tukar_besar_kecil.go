package CodeTest

import (
	"strings"
)

func tukarBesarKecil(huruf string) string{
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var output string
	arrhuruf := strings.Split(huruf, "")
	for _, s := range arrhuruf {
		if strings.Contains(key,s) {
			output += strings.ToLower(s)
		}else {
			output += strings.ToUpper(s)
		}
	}
	return output
}
