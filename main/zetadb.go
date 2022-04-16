package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	//IP address and port number of database server
	serverIpPort string = "127.0.0.1:40320"
	//"127.0.0.1:40320"
	//"153.92.210.106:40320"
)

func main() {
	//create a reader to read input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("log in or create a new user")

	for { //loop until quit is inserted
		fmt.Print("zetaDB> ")

		//read sql from user
		sqlBytes, _, err := reader.ReadLine()
		if len(sqlBytes) == 0 {
			continue
		}
		checkError(err)

		sql := string(sqlBytes)
		if strings.EqualFold(sql, "quit;") {
			os.Exit(0)
		}

		//assign server address
		tcp_addr, err := net.ResolveTCPAddr("tcp4", serverIpPort)
		checkError(err)
		// issue connection requirement
		conn, err := net.DialTCP("tcp", nil, tcp_addr)
		checkError(err)

		//socket read & write data
		_, err = conn.Write(sqlBytes)
		checkError(err)
		buffer := make([]byte, 16384)
		_, err = conn.Read(buffer)
		checkError(err)
		replyString := string(buffer)
		fmt.Println(string(replyString))

		// close connection
		conn.Close()
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

