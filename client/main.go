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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	check(err)

	addr := &syscall.SockaddrInet4{Port: 3000, Addr: [4]byte{0, 0, 0, 0}}
	err = syscall.Connect(fd, addr)
	check(err)

	_, err = syscall.Write(fd, []byte("hi\n"))
	check(err)

	data := make([]byte, 1024)
	_, err = syscall.Read(fd, data)
	check(err)
	fmt.Printf("Server sent: %s", string(data))
}
