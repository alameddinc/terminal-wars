package main

import (
	"fmt"
	. "github.com/alameddinc/terminal-wars/tasks/hunter"
)

var Connections []*Connection

func main() {
	//fmt.Println(repair.GetRepair())
	//fishing.GenerateMap()
	//routing.Handle()
	InitNetwork(&Connections)
	newConnection := Connection{}
	newConnection.Username = "Alameddin"
	newConnection.JoinNetwork(&Connections)
	fmt.Print("Started")
}
