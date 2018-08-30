package main

import (
	"flag"
	"fmt"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "ip for connect")
	var port int
	flag.IntVar(&port, "port", 8080, "port for connect")
	flag.Parse()
	fmt.Println(*ip, ":", port)
}
