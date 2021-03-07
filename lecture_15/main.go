package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//net = metwork
	nl, err := net.Listen("tcp", ":8888") // layer 2ta 1. network  2. Address(socket)ip & port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // normal exit er value 0(zero) / elligale exit 1(one)
	}

	conn, err := nl.Accept() // layer 5(session layer) session start kora- ekhane o same 2ta nite hove 1.connection(conn) & 2nd erro(err)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // ekhane o exit nivo coz ekaneo amar accecpt e hoi ni tai
	}

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr) //[::1]:14995
	// byte
	conn.Write([]byte("welcome we have recived your msg"))

	conn.Close()

}
