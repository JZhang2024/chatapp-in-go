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

func main() {
    // Listen on TCP port 8080 on all interfaces.
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error listening on port 8080: %v\n", err)
        os.Exit(1)
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080...")
    for {
        // Wait for a connection.
        conn, err := listener.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error accepting connection: %v\n", err)
            continue
        }
        // Handle the connection in a new goroutine.
        // The loop then returns to accepting, so that
        // multiple connections may be served concurrently.
        go handleConnection(conn)
    }
}