package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
)

type Account struct{ //Basic account structure. Name is a string, balance is a float32
	name string
	balance float32
}

var accounts map[int]Account //declare a map to manage all accounts

func setup() map[string]string{
	confileName := "/assets/config.csv"
	csvFile, err:= os.Open(confileName)
	for err != nil { //Handle config file opening error
		fmt.Println(err)
		fmt.Println("Please specify path to 'config.csv': ")
		_, err := fmt.Scan(&confileName)
		_, err := os.Open(confileName)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, err = reader.Read()

	configuration := make(map[string]string)
	for row, col := range(configfile){

	}
	return configuration
}

var loadAccountMap map(){
	var fileName string //Declare a string to hold the name of the textfile
	rawData, err := ioutil.ReadFile(fileName) //attempt to open the file
	if err != nil { //Handle file opening error
		fmt.Println(err)
		return
	}
	accounts = make(map[int]Account)
	accounts[165643] = Account{
		"Alice",40.68,
	}
	fmt.Println(accounts[165643])
}

func main() {
	configuration := setup()
	usertable := configuration(1,0)
	accounts := loadAccountMap(configuration(usertable))

}
