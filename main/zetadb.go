package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
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

	//variable to store current userId, -1 as initial value
	var currentUserId int32 = -1
	var currentUserName string = ""

	fmt.Println("log in or create a new user")

	for { //loop until quit is inserted
		if currentUserName == "" {
			fmt.Print("zetaDB> ")
		} else {
			fmt.Print("zetaDB@" + currentUserName + "> ")
		}

		//read sql from user
		sqlBytes, _, err := reader.ReadLine()
		if len(sqlBytes) == 0 {
			continue
		}
		checkError(err)

		//if user input is quit; exit immediately
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

		//create a new request
		currentRequest := NewRequest(currentUserId, sql)
		currentRequestBytes := currentRequest.RequestToBytes()

		//socket write
		_, err = conn.Write(currentRequestBytes)
		checkError(err)

		//socket read
		buffer := make([]byte, 16384)
		_, err = conn.Read(buffer)
		checkError(err)
		currentResponse := NewResponseFromBytes(buffer)

		//show response info
		if currentResponse.StateCode == 1 {
			logResponseMessages := strings.Split(currentResponse.Message, " ")
			currentUserIdINT, _ := strconv.Atoi(logResponseMessages[0])
			currentUserId = int32(currentUserIdINT)
			currentUserName = logResponseMessages[1]
			fmt.Println("log in successfully")
		} else {
			fmt.Println(currentResponse.Message)
		}

		// close connection
		conn.Close()

		//if system is stoped manually, exit this client
		if currentResponse.StateCode == -2 {
			os.Exit(0)
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
