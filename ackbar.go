package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	// Flag Setup
	port = flag.String("port", "", "Port to fuck over requests on")
)

func init() {
	flag.StringVar(port, "p", "", "Port to fuck over requests on")
}

func usage() {
	fmt.Println("ERROR: Missing argument.")
	fmt.Println("Usage:")
	fmt.Println("    ./ackbar -p 8080")
	os.Exit(1)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	log.Println("IT'S A TRAP!!")
	log.Println("REQUEST FROM:", conn.RemoteAddr())
	// OUR buffer is only 4096 bytes so that we are not vulnerable
	// to the thing we are doing to someone else.
	var buf [4096]byte
	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			log.Println("Got Error: ", err)
			log.Println("Closing Connection.")
			return
		}
		request := string(bytes.Trim(buf[:], "\x00"))
		log.Println("REQUEST CONTENT: ", request)
		log.Println("LOCKING S-FLOILS IN ATTACK POSITION")
		fbuf := make([]byte, 2048)
		n := 0
		for err == nil {
			i := 0
			_, err := rand.Read(fbuf)
			i, err = conn.Write(fbuf)
			n += i
			if err != nil {
				log.Println("TARGET DOWN:", err)
				break
			}
		}
		log.Printf("FIRED %d BYTES BEFORE ENDPOINT COWARDLY RETREATED.", n)
		return
	}
}

func main() {
	log.Println("Starting Ackbar...")
	flag.Parse()
	if len(*port) == 0 {
		usage()
	}
	log.Printf("Starting up listener on %s. Use ctrl-c to stop.", *port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+*port)
	checkError(err)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			// Don't use checkError because this shouldn't matter.
			log.Println("Listener encountered an error:", err)
		}
		go handleRequest(conn)
	}
}
