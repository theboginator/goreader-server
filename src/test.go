package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(conn, "679,64")

}
