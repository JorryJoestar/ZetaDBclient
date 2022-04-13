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
)

func main() {
	//create a reader to read input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("log in or create a new user")

	for { //loop until quit is inserted
		//assign server address
		tcp_addr, err := net.ResolveTCPAddr("tcp4", serverIpPort)
		checkError(err)
		// issue connection requirement
		conn, err := net.DialTCP("tcp", nil, tcp_addr)
		checkError(err)

		fmt.Print("zetaDB> ")

		sqlBytes, _, err := reader.ReadLine()
		checkError(err)

		sql := string(sqlBytes)
		if strings.EqualFold(sql, "quit;") {
			os.Exit(0)
		}

		//socket read & write data
		_, err = conn.Write(sqlBytes)
		checkError(err)
		buffer := make([]byte, 256)
		_, err = conn.Read(buffer)
		checkError(err)
		fmt.Println(string(buffer))

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
