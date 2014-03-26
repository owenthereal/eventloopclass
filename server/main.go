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

	addr := &syscall.SockaddrInet4{Port: 3000, Addr: [4]byte{0, 0, 0, 0}}
	err = syscall.Bind(acceptingFd, addr)
	check(err)

	err = syscall.Listen(acceptingFd, 100)
	check(err)

	for {
		connectionFd, _, err := syscall.Accept(acceptingFd)
		check(err)
		fmt.Println("Accepted new connectrion")

		err = syscall.SetNonblock(connectionFd, true)
		check(err)

		fdSet := &syscall.FdSet{Bits: [32]int32{int32(connectionFd)}}
		err = syscall.Select(connectionFd+1, fdSet, nil, nil, nil)
		check(err)

		data := make([]byte, 1024)
		_, err = syscall.Read(connectionFd, data)
		check(err)
		fmt.Printf("Received: %s\n", string(data))

		_, err = syscall.Write(connectionFd, data)
		check(err)

		err = syscall.Close(connectionFd)
		check(err)
	}
}
