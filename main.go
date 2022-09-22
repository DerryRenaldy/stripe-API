package main

import (
	"stripe-project/server"
)

func main() {

	sv := server.Register()

	sv.Start()

}
