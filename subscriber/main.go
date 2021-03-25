package main

import (
    "context"
    "log"
    "time"

    "github.com/aigent/nq"
)

func main() {
    opts := nq.SubOpts{
        KeepaliveTimeout: 5 * time.Second,
        Printf:           log.Printf,
    }
    sub := nq.NewSub("tcp4://:1234", opts, nq.NewDefaultMetrics())

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