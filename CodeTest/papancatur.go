package CodeTest

import "fmt"

func PapanCatur(ukuran int){
	for i := 0; i < ukuran; i++ {
		var baris = ""
		for j := 0; j < ukuran; j++ {
			if i %2 == 0 {
				baris += "# "
			} else {
				if j == ukuran -1 {
					baris += " "
				}else {
					baris += " #"
				}
			}
		}
		fmt.Println(baris)
	}
}
