package CodeTest

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T){
	hasil := PalindromeReverse("kodok")
	fmt.Println(hasil)
}

func TestPalindromeWoTemp(t *testing.T) {
	ress := PalindromeWoTemp("kodok")
	fmt.Println(ress)
}

func TestPalindromeRecursive(t *testing.T) {
	akhir := PalindromeRecursive("kodok", 0)
	fmt.Println(akhir)
}