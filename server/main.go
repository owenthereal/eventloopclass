package main

import (
	"fmt"
	"log"
	"syscall"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	acceptingFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	check(err)

	err = syscall.Bind(acceptingFd, &syscall.SockaddrInet4{Port: 3000, Addr: [4]byte{0, 0, 0, 0}})
	check(err)

	err = syscall.Listen(acceptingFd, 100)
	check(err)

	for {
		connectionFd, _, err := syscall.Accept(acceptingFd)
		check(err)
		fmt.Println("Accepted new connectrion")

		data := make([]byte, 1024)
		_, err = syscall.Read(connectionFd, data)
		check(err)
		fmt.Printf("Received: %s\n", string(data))

		syscall.Write(connectionFd, data)
		syscall.Close(connectionFd)
	}
}
