package main

import "fmt"

func main() {
	dict := map[string]int{"one": 1, "two": 2}
	keys := make([]string, 0, len(dict))
	values := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
		values = append(values, dict[k])
	}

	fmt.Print(keys)
	fmt.Print(values)
}
