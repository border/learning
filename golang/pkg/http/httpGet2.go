package main

import (
    "log"
    "http"
    "net"
    "io/ioutil"
)

func main() {

    url, err := http.ParseURL("http://bbs.golang-china.org/")

    if err != nil {
        log.Exit(err)
    }

    tcpConn, err := net.Dial("tcp", "", url.Host+":80")

    if err != nil {
        log.Exit(err)
    }

    clientConn := http.NewClientConn(tcpConn, nil)

    var req http.Request
    req.URL = url
    req.Method = "GET"
    req.UserAgent = "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.11 (KHTML, like Gecko) Chrome/9.0.570.0 Safari/534.11"
    req.Header = map[string]string{}
    req.Header["Connection"] = "keep-alive"

    err = clientConn.Write(&req)
    if err != nil {
        log.Exit(err)
    }
    resp, err := clientConn.Read()
    if err != nil {
        log.Exit(err)
    }
    defer resp.Body.Close()

    log.Println("Http Response: " + resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)

    log.Println(string(body))
}
