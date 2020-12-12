package main

import (
	"fmt"
	"github.com/alameddinc/terminal-wars/tasks/repair"
)

func main() {
	fmt.Println(repair.CheckRepair(15765233, "2f0899f61805126a241830dc8f0b3d70"))
	fmt.Println(repair.CheckRepair(64176233, "d723f8c7e136c804c32aa19236dba4d3"))
	fmt.Println(repair.CheckRepair(32165233, "4407ddc7463403b98f25ae796bba07f8"))
	//fmt.Println(repair.GetRepair())
	//fishing.GenerateMap()
	//routing.Handle()
}
