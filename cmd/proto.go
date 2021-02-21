package main

import (
    "fmt"
    "flag"
    "net"
    "strings"
    "io/ioutil"
    "log"

    "github.com/golang/protobuf/proto"
	"github.com/yaowenqiang/moderngo/protobuf"
)

func main() {
    op := flag.String("op", "o", "s for server, c for lcient")
    flag.Parse()
    switch strings.ToLower(*op) {
    case "s":
        RunProtoServer()
    case "c":
        RunProtoClient()
    }
}

func RunProtoServer() {
    l, err := net.Listen("tcp", "127.0.0.1:8282")
    if err != nil {
        log.Fatal(err)
    }

    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }

        defer l.Close()
        go func(c net.Conn) {
            defer c.Close()
            data, err := ioutil.ReadAll(c)

            if err != nil {
                return
            }

            a := &protobuf.Animal{}

            err = proto.Unmarshal(data, a)
            if err != nil {
                log.Println(err)
                return
            }
            fmt.Println(a)
        }(c)
    }
}


func RunProtoClient() {
    a := &protobuf.Animal{
        Id: 1,
        AnimalType:"dino",
        Nickname: "jack",
        Zone: 1,
        Age: 10,
    }

    data, err := proto.Marshal(a)
    if err != nil {
        log.Fatal(err)
    }

    SendData(data)

}

func SendData(data []byte) {
    c, err := net.Dial("tcp", "127.0.0.1:8282")
    if err != nil {
        log.Fatal(err)
    }

    defer c.Close()
    c.Write(data)
}
