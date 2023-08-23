package main

import (
	shifter "github.com/TeamTutx/pg-shifter"
	"github.com/TeamTutx/plib/database/postgresql"
)

func main() {
	if conn, err := postgresql.Conn(true); err == nil {
		//this will create the test_address go struct in filepath
		//as filepath is not given so it will be created in
		//pwd/log/TestAddress/TestAddress.go
		shifter.NewShifter().CreateStruct(conn, "test_address", "")
	}
}
