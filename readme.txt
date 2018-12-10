GoReader - A server for listening to a RFID-reading client running the GoReader Client Software

Settings stored in config.csv within the '/assets' directory

Program functions:
-Listens to a a reader sending data to a configurable network port, receives an id # and a transaction amount from reader
-Looks up id # on table and adds/subtracts the transaction amount from the principal balance
-Will decline the transaction if the transaction would lower the principal balance below $0 OR the account is locked
-Reports back to reader application "PASS/FAIL"
-Allows intuitive editing of user table and balances
-Allows editing of network parameters
