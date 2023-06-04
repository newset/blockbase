package main

import (
	tron "blockbase/pkg/tron"
)

func main() {
	account, _ := tron.CreateNewTronAccount()
	println("addres: ", account)
}
