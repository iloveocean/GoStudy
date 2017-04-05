package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Home struct {
	ID   bson.ObjectId `bson:"_id, omitempty"`
	Name string        `bason:"name"`
}

func main() {
	h := Home{Name: "Peter"}
	raw := []byte(h.ID)
	fmt.Printf("raw length is: %d\n", len(raw))
	for index, item := range raw {
		fmt.Printf("index %d value is: %d\n", index, item)
	}
}
