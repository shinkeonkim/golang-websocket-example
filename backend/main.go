package main

import "github.com/shinkeonkim/golang-websocket-example/network"

func main() {
	n := network.NewServer()
	n.StartServer()
}
