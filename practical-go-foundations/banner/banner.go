package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("hello", 8)
	banner("😁", 8)
	// code point = rune ~= unicode character

	s := "G😃"
	for i, r := range s {
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	i := isPalindrome("geeeg")

	fmt.Println(i)

	// byte (uint8)
	// rune (int32)
	// %c is unicode representation?
	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	fmt.Printf("%20s\n", s)
}

// bad implx
// func isPalindrome(s string) bool {
// 	tmp := []byte{}
// 	for i := range s {
// 		tmp = append([]byte{s[i]}, tmp...)
// 	}
// 	return string(tmp) == s
// }

func isPalindrome(s string) bool {
	for i := 0; i < (len(s) / 2); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2
	// RuneCountInString fixes the misalignment for strings
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
