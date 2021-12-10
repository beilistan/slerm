package main

import (
	"fmt"
	"strings"
)

func charCount(sample string) {
	m := make(map[string]int)
	for _, letter := range sample {
		c := strings.Count(sample, string(letter))
		m[string(letter)] = c
	}
	for k, v := range m {
		fmt.Printf("Char:%s Count: %d \n", k, v)
	}
}

func main() {
	sample := "съешь ещё этих мягких французских булок, да выпей чаю"
	charCount(sample)
}
