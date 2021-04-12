package main

import (
	"crypto/tls"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	//    conf := &tls.Config{
	//InsecureSkipVerify: true,
	//    }

     // openssl genrsa -out server.key 2048
     //openssl ecparam -genkey -name secp384r1 -out server.key 
     //openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
  
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}
	conf := &tls.Config{Certificates: []tls.Certificate{cer}}

	conn, err := tls.Dial("tcp", "sunucu.server.lab:8090", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("merhaba\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
