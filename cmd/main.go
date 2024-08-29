package main

import (
	"githb.com/steelthedev/go-solana-todo/pkg/server"
	"githb.com/steelthedev/go-solana-todo/pkg/solana"
)

func main() {

	// Initiate solana client globally
	solana.InitClient()

	// Initiate server
	(server.StartServer(":3000"))
}
