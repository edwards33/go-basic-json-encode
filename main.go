package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Firstname string
	Lastname string
	Age int
}

func main() {
	pFirst := person{"Ed", "Navi", 20}
	json.NewEncoder(os.Stdout).Encode(pFirst)
}
