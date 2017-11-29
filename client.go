// client side
// after running the command on server.go
// compile with command in a seperate terminal: 
// $ go run client.go

package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {

	// connect to given socket (standard)
	conn, _ := net.Dial("tcp", "127.0.0.1:8081") 

	for {
		
		// listen for reply from server
		message, _ := bufio.NewReader(conn).ReadString('\n')
		
		// print the message sent from server
		fmt.Print("Message from server: " + message + "\n")
	}
}
