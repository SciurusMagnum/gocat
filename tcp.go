package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn       net.TCPConn
	inChannel  chan string
	connChanel chan string
	inReader   bufio.Reader
	conReader  bufio.Reader
	connWriter bufio.Writer
}

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

func (client *Client) start() {
	go readBuffer(client.connReader, client.connChannel)
	go readBuffer(client.ioReader, client.inChannel)
	go printChannelData(client.connChannel)
}

func newClient(ip string, port uint8) *Client {
	servAddr := fmt.Sprintf("%s:%d", ip, port)
	conn := initializeConnection(servAddr)
	// defer conn.Close()
	ioReader := bufio.NewReader(os.Stdin)

	connReader := bufio.NewReader(conn)
	client := &Client{

		inChannel:   make(chan string),
		connChannel: make(chan string),
		inReader:    bufio.NewReader(os.Stdin),
		connReader:  bufio.NewReader(conn),
	}
	client.start()
	return client

}
