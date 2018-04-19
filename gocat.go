package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func readBuffer(reader *bufio.Reader, channel chan string) {
	for {
		chanString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading from reader")
			close(channel)
			os.Exit(1)
		}
		chanString = strings.Trim(chanString, "\n")
		channel <- chanString

	}
}

func printChannelData(channel chan string) {
	for data := range channel {
		println(data)
	}
}
func sendChannelData(channel chan string, conn *net.TCPConn) {
	writer := bufio.NewWriter(conn)
	for data := range channel {
		_, err := writer.WriteString(data + "\n")
		if err == nil {
			err = writer.Flush()
		} else {
			println("connection closed")
			break
		}
	}
}

func initializeConnection(servAddr string) *net.TCPConn {

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
	return conn
}

func main() {
	servAddr := "localhost:6666"
	conn := initializeConnection(servAddr)
	defer conn.Close()
	ioReader := bufio.NewReader(os.Stdin)
	inChannel := make(chan string)

	connReader := bufio.NewReader(conn)
	connChannel := make(chan string)
	go readBuffer(connReader, connChannel)
	go readBuffer(ioReader, inChannel)
	go printChannelData(connChannel)

	sendChannelData(inChannel, conn)
}
