package main

import (
    "fmt"
    "net"
 )
 
 func main() {
 
    fmt.Println("[*] Connecting to Target")
    
    conn, err := net.Dial("tcp", "192.168.56.101:9999")
    if err != nil {
        fmt.Println("[!] There was an error connecting")
        fmt.Println("=====================================")
        fmt.Println(err)
        return
    }
    
    fmt.Fprintf(conn, "EXIT")
    conn.Close()
    
 
 }
