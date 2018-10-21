package main

import (
	"bytes"
	"fmt"
)

func main() {
	byt := bytes.NewBufferString("abcd123")
	fmt.Println(byt.Bytes())

	fmt.Println(byt.Len())
}
