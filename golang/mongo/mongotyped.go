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

func main() {
    conn, _ := mongo.Connect("127.0.0.1")
    coll := conn.GetDB("example").GetCollection("typed_structs")

    firstDoc := &ExampleDoc{
        Title: "Page Title",
        Body:  "Some example text",
        Comments: []Comment{
            Comment{
                Author:  "foobar",
                Message: "Thanks For the post",
            },
            Comment{
                Author:  "hokapok",
                Web:     "go.hokapok.com",
                Message: "No Problem",
            },
            Comment{
                Author:  "golang",
                Web:     "golang.org",
                Message: "Thanks",
            },
        },
    }

    bsonDocIn, _ := mongo.Marshal(firstDoc)
    coll.Insert(bsonDocIn)

    fmt.Println("Inserted Document")
    coll.Drop()
}
