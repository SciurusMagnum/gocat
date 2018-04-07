package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// strEcho := "Hello\n"
	servAddr := "localhost:6666"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed :", err.Error())
		os.Exit(1)
	}
	// println("write to server = ", strEcho)
	ioReader := bufio.NewReader(os.Stdin)
	// send to socket
	// listen for reply
	connReader := bufio.NewReader(conn)
	for {
		text, _ := ioReader.ReadString('\n')
		fmt.Fprintf(conn, text)
		// _, err = conn.Write([]byte(strEcho))
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}
		// reply := make([]byte, 1024)
		reply, err := connReader.ReadString('\n')
		reply = strings.Trim(reply, "\n	")
		// _, err = conn.Read(reply)
		if err != nil {
			println("read from server failed:", err.Error())
			os.Exit(1)
		}

		println(string(reply))

	}
	// conn.Close()
}
