package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}


}

func handle(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		//fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[1]
			fmt.Println("URL", m)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}