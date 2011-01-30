package main

import (
    "log"
    "io/ioutil"
    "http"
)

func main() {
    res, _, err := http.Get("http://bbs.golang-china.org/")
    if err != nil {
        log.Exit(err)
    }

    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)

    log.Println(string(body))
}
