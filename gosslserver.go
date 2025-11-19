package main

import (
    "log"
    "crypto/tls"
    "net"
    "bufio"
    "math/rand"
    "time"
    "fmt"
)

func main() {
    log.SetFlags(log.Lshortfile)
     //openssl genrsa -out server.key 2048
     //openssl ecparam -genkey -name secp384r1 -out server.key 
     //openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
    cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        log.Println(err)
        return
    }

    config := &tls.Config{Certificates: []tls.Certificate{cer}}
    ln, err := tls.Listen("tcp", ":8090", config) 
    if err != nil {
        log.Println(err)
        return
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    for {
        msg, err := r.ReadString('\n')
        if err != nil {
            log.Println(err)
            return
        }

        println(msg)

        // Random mesaj gönder
        randomMsg := generateRandomMessage()
        n, err := conn.Write([]byte(randomMsg + "\n"))
        if err != nil {
            log.Println(n, err)
            return
        }
    }
}

// Random sayı üretir (0-100 arası)
func generateRandomNumber() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(101)
}

// Random mesaj üretir
func generateRandomMessage() string {
    messages := []string{
        "Merhaba",
        "Hello World",
        "Selam",
        "Hi there",
        "Nasılsın",
    }
    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(messages))
    randomNum := generateRandomNumber()
    return fmt.Sprintf("%s - Random: %d", messages[randomIndex], randomNum)
}
