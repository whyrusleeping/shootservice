package main

import (
	"net"
	"fmt"
	"net/http"
)

var DoShoot chan bool

func WaitForReader(port string) {
	list, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	for {
		con, err := list.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection!")
		for {
			<-DoShoot
			_,err := con.Write([]byte("SHOOT\n"))
			fmt.Println("Sent Shot!")
			if err != nil {
				//Buffer another shoot command
				go func() {DoShoot <- true}()
				fmt.Println(err)
			}
		}
	}
}

func main() {
	DoShoot = make(chan bool)
	go WaitForReader(":8001")
	http.HandleFunc("/shoot", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received shoot request.")
		DoShoot <- true
	})

	panic(http.ListenAndServe(":8000", nil))
}
