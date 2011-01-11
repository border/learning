package main
/*
   http://go.hokapoka.com/golang/mongodb-golang-gomongo/ 
 */

import (
	"github.com/mikejs/gomongo/mongo"
	"fmt"
)

func main() {
	conn, _ := mongo.Connect("127.0.0.1")
	coll := conn.GetDB("example").GetCollection("map_string")

	firstDoc := map[string]string{
		"_id":   "1",
		"title": "Page Title",
		"body":  "Some example Text",
	}

	bsonDocIn, _ := mongo.Marshal(firstDoc)
	coll.Insert(bsonDocIn)

	fmt.Println("Inserted Document")
	coll.Drop()
}
