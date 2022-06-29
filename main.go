package main

import (
	_ "github.com/lib/pq"
	db "servermanagement.com/connection"
	h "servermanagement.com/handler"
)

func main() {
	db.Connect()
	h.HandleFunc()

}
