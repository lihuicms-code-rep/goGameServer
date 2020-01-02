package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("client connect to server error ", err)
		return
	}

	fmt.Println("[Client] connect to server localhost:8888 success")

	for {
		if _, err := conn.Write([]byte("i am client msg")); err != nil {
			fmt.Println("client write msg to server error ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client receive data from server error ", err)
			return
		}

		fmt.Println("client receive server callback data ", string(buf[:cnt]))

		time.Sleep(1 * time.Second)
	}
}
