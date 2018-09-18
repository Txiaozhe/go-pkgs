package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	i, _ := strconv.Atoi("123")
	fmt.Println(i)

	i64, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(i64)

	str := strconv.Itoa(123)
	fmt.Println(str)

	str64 := strconv.FormatInt(1234, 10)
	fmt.Println(str64)

	idx := strings.Index("h", "hello")
	fmt.Println(idx)

	join_str := strings.Join([]string{"he", "is"}, ",,,")
	fmt.Println(join_str)
}
