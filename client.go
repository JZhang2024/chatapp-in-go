package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        os.Exit(1)
    }
    defer conn.Close() // Ensure connection is closed on function exit

    // Create a reader for stdin once instead of in the loop
    reader := bufio.NewReader(os.Stdin)

    // Create a reader for the connection once instead of in the loop
    connReader := bufio.NewReader(conn)

    for {
        fmt.Print("Enter message: ")
        text, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from stdin:", err)
            continue // Skip this iteration
        }

        // Send the text to the server
        _, err = fmt.Fprintf(conn, text+"\n")
        if err != nil {
            fmt.Println("Error sending message to server:", err)
            continue // Skip this iteration
        }

        // Read the message from the server
        message, err := connReader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading message from server:", err)
            continue // Skip this iteration
        }
        fmt.Print("Message from server: " + message)
    }
}