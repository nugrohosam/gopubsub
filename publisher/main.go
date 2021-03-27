package main

import (
    "context"
    "log"
    "time"
    "strconv"
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

    opts := nq.PubOpts{
        KeepaliveTimeout: 5 * time.Second,
        ConnectTimeout:   3 * time.Second,
        WriteTimeout:     3 * time.Second,
        FlushFrequency:   100 * time.Millisecond,
        NoDelay:          true,
        Printf:           log.Printf,
    }
    pub := nq.NewPub("tcp4://"+*listenHost+":"+*listenPort, opts, nq.NewDefaultMetrics())
    for {
        // Publish the message using 100 connections
        for i := 1; i <= 100; i++ {
            ke := "Hello nanoQ ke-" + strconv.Itoa(i+1)
            if err := pub.Publish(context.TODO(), []byte(ke), i); err != nil {
                log.Println("Error while publishing:", err)
            }
        }
    }
}