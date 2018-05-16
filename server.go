package main

import (
	"net"
	"fmt"
	"time"
	"strconv"
)

var num_of_clients = 0

// go routine function to send time to client
func SendTime(conn net.Conn) {

	// run the loop until cmd-c is pressed
	for {

		// get current time
		t:=time.Now()

		// format the time for string
		time_string := strconv.Itoa(t.Hour() % 13 + 1) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())

		// sends string to client
		ip_address := conn.LocalAddr().String() 

		// write the message including time, ip, number of clients
		conn.Write([]byte("Time: " + time_string + "\tIP: " + ip_address + "\tNumber Of Clients: " + strconv.Itoa(num_of_clients) + "\n"))
		
		// sleep for 1 second and update
		time.Sleep(1000 * time.Millisecond) 
	}
}

func main() {

	fmt.Println("Launching server...")

	// call listen for all interfaces
	ln, _ := net.Listen("tcp", ":8081") 

	for {
		
		fmt.Println("Waiting for client connection...")

		// accepting connection on port
		conn, _ := ln.Accept() 

		// once accepted, new client is connected
		fmt.Println("Connected to new client!")

		// incrememnt every time a new client is connected
		num_of_clients++

		// each client gets their own go routine
		go SendTime(conn) 
	}
}
