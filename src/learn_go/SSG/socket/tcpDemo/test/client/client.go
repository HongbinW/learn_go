package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn,err := net.Dial("tcp","0.0.0.0:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("client dial err:",err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client read from stdin err :", err)
		}
		if line == "exit\n" {
			break
		}
		conn.Write([]byte(line))
	}
}
