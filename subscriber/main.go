package main

import (
    "context"
    "log"
    "time"
    "fmt"
    "flag"

    "github.com/aigent/nq"
)

func main() {
	listenHost := flag.String("host", "none", "--")
	listenPort := flag.String("port", "none", "--")
	flag.Parse()

    if *listenHost == "none" {
        fmt.Println("flag [--host] not defined")
        return
    }
    
    if *listenPort == "none" {
        fmt.Println("flag [--port] not defined")
        return
    }

    opts := nq.SubOpts{
        KeepaliveTimeout: 5 * time.Second,
        Printf:           log.Printf,
    }

    sub := nq.NewSub("tcp4://"+*listenHost+":"+*listenPort, opts, nq.NewDefaultMetrics())
    go func() {
        buf := make([]byte, 4096)
        for {
            if msg, stream, err := sub.Receive(context.TODO(), buf); err != nil {
                log.Println("Error while receiving:", err)
                continue
            } else {
                log.Printf("message from stream '%v' is: %s\n", stream, msg)
            }
        }
    }()
    if err := sub.Listen(context.TODO()); err != nil {
        log.Println("Listen error:", err)
    }
}