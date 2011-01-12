package main

import (
    "github.com/mikejs/gomongo/mongo"
    "fmt"
)

type ExampleDoc struct {
    Title    string
    Body     string
    Comments []Comment
}

type Comment struct {
    Author  string
    Web     string
    Message string
}

type searchDoc struct {
    title string
}

func main() {
    conn, _ := mongo.Connect("127.0.0.1")
    coll := conn.GetDB("example").GetCollection("typed_structs")

    qFindDoc, _ := mongo.Marshal(&searchDoc{title: "Page Title"})
    bsonDocOut, _ := coll.FindOne(qFindDoc)

    var foundDoc ExampleDoc
    mongo.Unmarshal(bsonDocOut.Bytes(), &foundDoc)

    fmt.Println(foundDoc.Title + "::" + foundDoc.Body)

    fmt.Printf("Comment %v\n", foundDoc)

    coll.Drop()
}
