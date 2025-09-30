package main

import (
	"fmt"

	"compiladores-tarea1-go/pkg/ds/hash"
)

func main() {
	hashTable := hash.Init()
	list := []string{"ERIC", "KENNY", "KYLE", "STAN", "RANDY", "BUTTERS"}
	for _, v := range list {
		hashTable.Insert(v)
	}
	fmt.Println(hashTable)
	fmt.Println(hashTable.Search("RANDY"))

	hashTable.Delete("RANDY")
	fmt.Println(hashTable.Search("RANDY"))
}
