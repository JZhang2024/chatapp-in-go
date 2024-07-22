package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading from connection: %v\n", err)
            break // Exit the loop if an error occurs
        }
        fmt.Printf("Message Received: %s", message)
        
        _, err = conn.Write([]byte(message)) // Echo back the message to the client
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error writing to connection: %v\n", err)
            break // Exit the loop if an error occurs
        }
    }
}