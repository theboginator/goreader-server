package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

var (
	outputChan   = make(chan string)       //Channel for accepted/denied message going to PED
	idChan       = make(chan int64)        // Channel for incoming PED id scans
	valueChan    = make(chan float64)      //Channel for transaction value
	accounts     = make(map[int64]float64) //Create the user account table
	defaultValue = 150.00                  //Default new account balance
)

func configureConnections(conn net.Conn) { //Setup connection manager to handle incoming/outgoing data
	go idReader(conn)
	go valReader(conn)

}

func clientWriter(conn net.Conn, result string) { //Send a "accepted/declined" message in the form of a boolean
	fmt.Fprintln(conn, result)
}

func idReader(conn net.Conn) { //Read userID sent from Pi
	input := bufio.NewScanner(conn)
	for input.Scan() {
		data := input.Text()
		id, _ := strconv.ParseInt(data, 10, 64)
		idChan <- id
	}
}

func valReader(conn net.Conn) { //Read transaction amount sent from Pi
	input := bufio.NewScanner(conn)
	for input.Scan() {
		data := input.Text()
		value, _ := strconv.ParseFloat(data, 64)
		valueChan <- value
	}
}

func handleTransaction(id int64, value float64) bool { //Attempt to process the transaction
	var accept bool
	if _, exists := accounts[id]; !exists { //check to see if the id exists in the table, if not, create it.
		accounts[id] = defaultValue //Give it a default value
	}
	balance := accounts[id]        //Retrieve the balance of the account
	tempbalance := balance - value //test out the computation
	if tempbalance < 0 {
		accept = false //Reject transaction
	} else {
		accounts[id] = tempbalance //Update the new balance
		accept = true              //Accept transaction
	}
	updateTable()
	return accept
}

func updateTable() {
	for id, balance := range accounts {
		fmt.Printf("key[%s] value[%s]\n", id, balance)
	}
}

func main() {
	networkMgr, err := net.Listen("tcp", "localhost:8000") //setup connection
	if err != nil {
		log.Fatal(err) //handle an error
	}

	for {
		conn, err := networkMgr.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		configureConnections(conn)
		select {
		case newID := <-idChan:
			reqValue := <-valueChan
			accepted := handleTransaction(newID, reqValue)
			if accepted {
				clientWriter(conn, "Approved")
			} else {
				clientWriter(conn, "Declined")
			}
		}
	}
}
