package main

import (
	"fmt"
	"strings"
)

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
// Contoh : Input  : “aaasssiia”output  : a: 4, s: 3, i: 2

func SimpleFunction(str string) {
	strs := map[string]int{}

	for _, v := range strings.Split(str, "") {
		if strs[v] == 0 {
			strs[v]++
		} else {
			strs[v] += 1
		}
	}

	fmt.Println(strs)
}
// Output : map[a:8 b:3 e:1 f:4]
