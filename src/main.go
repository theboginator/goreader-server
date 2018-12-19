package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

var (
	outputChan   = make(chan string)       //Channel for accepted/denied message going to PED
	idChan       = make(chan int64)        // Channel for incoming PED id scans
	valueChan    = make(chan float64)      //Channel for transaction value
	accounts     = make(map[int64]float64) //Create the user account table
	defaultValue = 150.00                  //Default new account balance
)

func configureConnections(conn net.Conn) { //Setup connection manager to handle incoming/outgoing data
	go cardReader(conn)
}

func clientWriter(conn net.Conn, result string) { //Send a "accepted/declined" message in the form of a boolean
	fmt.Fprintln(conn, result)
}

func cardReader(conn net.Conn) { //Read userID sent from Pi
	input := bufio.NewScanner(conn)
	for input.Scan() {
		raw := input.Text()                         //get raw data
		var data = strings.Split(raw, ",")          //Split it into the userid and transaction amount
		id, _ := strconv.ParseInt(data[0], 10, 64)  //userid
		value, _ := strconv.ParseFloat(data[1], 64) //transaction amount
		idChan <- id                                //load the userid into the IDChan
		valueChan <- value                          //load the value into the ValueChan
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
	return accept //Return whether the transaction succeeded or not
}

func updateTable() { //Update the account map
	for id, balance := range accounts {
		fmt.Printf("User ID: [%s] Balance: [%s]\n", id, balance)
	}
}

func main() {
	networkMgr, err := net.Listen("tcp", "localhost:6000") //setup connection
	if err != nil {
		log.Fatal(err) //handle an error
	}

	for {
		conn, err := networkMgr.Accept() //Accept a connection from the PED
		if err != nil {
			log.Print(err)
			continue
		}
		configureConnections(conn)                     //Setup the goroutines that will listen for data and put it in the appropriate channels
		newID := <-idChan                              //Once the client sends a transaction, record the account number
		reqValue := <-valueChan                        //Record the transaction amount
		accepted := handleTransaction(newID, reqValue) //Attempt the transaction
		if accepted {
			clientWriter(conn, "Approved") //Send an approved message
		} else {
			clientWriter(conn, "Declined") //Send a declined message
		}
	}
}
