package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {

	p1 := person{
		First: "Joe",
	}

	p2 := person{
		First: "Peter",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)

	if err != nil {
		log.Panic(err)

	}
	fmt.Println("Print JSON", string(bs))

	xp2 := []person{}
	err = json.Unmarshal(bs, &xp2)

	if err != nil {
		log.Panic(err)

	}

	fmt.Println("Convert back to Go structure", xp2)

}