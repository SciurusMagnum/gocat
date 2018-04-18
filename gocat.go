package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func readBuffer(reader *bufio.Reader, channel chan string) {
	for {
		var text []byte
		var chanString string
		reader.Read(text)
		chanString = strings.Trim(string(text), "\n	")
		if chanString != "" {
			channel <- chanString
		}

	}
	// fmt.Fprintf(conn, text)
}

// if err != nil {
// 				println("read from server failed:", err.Error())
// 				os.Exit(1)
// 			}
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
	inChannel := make(chan string)
	// send to socket
	// listen for reply
	connReader := bufio.NewReader(conn)
	connChannel := make(chan string)
	go readBuffer(connReader, connChannel)
	go readBuffer(ioReader, inChannel)

	for {
		// var inText, connText string
		// 	text, _ := ioReader.ReadString('\n')
		// 	fmt.Fprintf(conn, text)
		// reply := make([]byte, 1024)
		select {
		case <-connChannel:
			println(<-connChannel)
		case <-inChannel:
			fmt.Fprintf(conn, <-inChannel)
		default:
			time.Sleep(50 * time.Millisecond)

		}

	}
	// reply, err := connReader.ReadString('\n')
	// reply = strings.Trim(reply, "\n	")
	// // _, err = conn.Read(reply)
	// if err != nil {
	// 	println("read from server failed:", err.Error())
	// 	os.Exit(1)
	// }

	// println(string(reply))

	// conn.Close()
}
