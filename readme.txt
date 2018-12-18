GoReader - A server for listening to a RFID-reading client running the GoReader Client Software
Written by Jacob Bogner
Current Version: 0.5

Settings stored in config.csv within the '/assets' directory

Program functions:
- Listens to a network port for incoming messages from a PED (Payment Entry Device - a Raspberry Pi running the
goreader-client application)

- Manages a map of all user accounts and their associated balances

- When PED reports a new account number and transaction amount, the number is looked up in the map and the transaction
amount is added/subtracted from the balance. If the new balance is <0, transaction will be rejected and a "denied"
message will be sent to PED, else a "success" message will be sent to PED and the account balance will be updated

To Install:

Open GoLand and check out the latest version from https://github.com/theboginator/goreader-server.git
go build main.go
go run main.go