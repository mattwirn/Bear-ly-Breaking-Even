package main

import (
	"encoding/json"
	"fmt"
)

type Ex1 struct {
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
}
type Ex2 struct {
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
}

func main() {
	first := Ex1{
		Name:   "bob",
		Amount: 69,
	}
	second := Ex2{
		Name:   "Rob",
		Amount: 96,
	}
	all := []interface{}{
		first,
	}
	a, _ := json.Marshal(all)
	fmt.Println(string(a))

	all = append(all, second)
	a, _ = json.Marshal(all)
	fmt.Println(string(a))
}
