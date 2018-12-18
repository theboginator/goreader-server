package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Account struct { //Basic account structure. Name is a string, balance is a float32
	name    string
	balance float32
}

var (
	outputChan = make(chan bool)   //Channel for messages going to PED
	inputChan  = make(chan string) // Channel for incoming PED messages
	accounts   map[int]Account     //declare a map to manage all accounts
)

func configureConnections(conn net.Conn) {
	output := make(chan string, 10) // outgoing client data
	go clientWriter(conn, output)
	in := make(chan string) // incoming client data
	go clientReader(conn, in)
}

func clientWriter(conn net.Conn, ch <-chan bool) {
	fmt.Fprintln(conn, message)
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

func main() {
	networkMgr, err := net.Listen("tcp", "localhost:8000") //setup connection
	if err != nil {
		log.Fatal(err) //handle an error
	}

	//go sendMessage()
	for {
		conn, err := networkMgr.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go configureConnections(conn)

	}
}
