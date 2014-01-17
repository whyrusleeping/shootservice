package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must specify server!")
	}
	server := os.Args[1]
	con, err := net.Dial("tcp", server)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(con)
	for scan.Scan() {
		if scan.Text() == "SHOOT" {
			//Fire?
			fmt.Println("Go shoot someone.")
		}
	}
}

