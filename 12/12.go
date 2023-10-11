package main

import "fmt"

func main() {
	arr := [...]string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, str := range &arr {
		set[str] = struct{}{}
	}
	fmt.Println(set)
}
