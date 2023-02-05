package main

import (
	"fmt"
)
type people struct{
	name string
	salary int
	height float32
	weight float32
}
func main() {
	name := []string{"a", "b", "c", "d", "f", "r", "s"}
	fmt.Println(len(name))
	fmt.Println(name)
	name = append(name, "he")
	fmt.Println(name[:])
	if 10 == 2 {
		fmt.Println("True")
	} else {
		fmt.Println("false")
	}
	number := 2
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("else")
	}
	code := map[string]string{"a": "ant", "c": "cat", "d": "dog", "e": "eyes"}

	for key, value := range code {
		fmt.Println("value = ", key, value)
	}
	fmt.Println(ans(3))

	p1 := people{name: "A", salary: 20000, height: 165.2,weight: 82.7}
	fmt.Println(p1)
}
func ans(i int) (int, int) {
	x := 2
	return x + i, i
}
