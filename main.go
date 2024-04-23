package main

import (
	"fmt"
	"strings"
)

type Response struct {
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

func main() {
	var str string
	fmt.Scan(&str)

	CountUniqueChar(str)
}

func CountUniqueChar(Str string) {
	arr := strings.Split(Str, "")
	counts := make(map[string]int)

	// Count the occurrences
	for _, str := range arr {
		for _, char := range str {
			counts[string(char)]++
		}
	}

	// Print the counts
	for char, count := range counts {
		fmt.Printf("%s: %d\n", char, count)
	}

}

// Buat function untuk hitung jumlah masing-masing karakter unik yang berurutan dari sebuah string.Fn Name : CountUniqueChar(input)Input  : StringOutput  : []map[string]int
// Contoh : Input  : “aaasssiia”output  : [[“a”: 4], [“s”: 3], [“i”: 2]]
