package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ReverseString("Gello")
	fmt.Println(s)
}

func ReverseString(s string) string {
	if s != "" {
		splitString := strings.Split(s, "")

		for i, j := len(s)-1, 0; j <= i; i, j = i-1, j+1 {
			splitString[i], splitString[j] = splitString[j], splitString[i]
		}
		return strings.Join(splitString, "")
	}

	return ""
}
