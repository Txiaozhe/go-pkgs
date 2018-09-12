package main

import (
	"fmt"
)

var arr1 = [8]string{"go", "python", "java", "c", "c++", "php"}
var arr2 = [8]string{0: "go", 1: "python", 2: "java", 3: "c", 4: "c++", 5: "php"}
var arr3 = [...]string{"go", "python", "java", "c", "c++", "php"}

func main() {
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println(len(arr3))
}
