package CodeTest

import "strings"

func UbahHuruf(huruf string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	var output string
	str := strings.Split(huruf, "")
	for _, v := range str {
		pos := strings.Index(alphabet, v)
		output += alphabet[pos+1:pos+2]
	}
	return output
}
