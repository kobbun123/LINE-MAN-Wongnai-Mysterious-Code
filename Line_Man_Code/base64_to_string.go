package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

var valid = map[rune]bool{}

func init() {
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=" {
		valid[r] = true
	}
}

func clean(s string) string {
	return strings.Map(func(r rune) rune {
		if valid[r] {
			return r
		}
		return -1
	}, s)
}

func ReverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
	}
	return string(runes)
}

func main() {
	s := `aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K`
	s = clean(s)
	dcd, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dcd))
	fmt.Println(ReverseString(string(dcd)))
	result := strings.ReplaceAll(ReverseString(string(dcd)), ":", " ")
	fmt.Println(result)

}
