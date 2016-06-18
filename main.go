package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func main() {

	//path := os.Getenv("PATH")

	user, err := user.Lookup(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gid, err := strconv.Atoi(user.Gid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(pw.Uid)
	syscall.Setuid(uid)
	syscall.Setgid(gid)
	if strings.HasPrefix(os.Args[2], "/") {
		err := syscall.Exec(os.Args[2], os.Args[2:], os.Environ())
		if err != nil {
			fmt.Println(err)
		}
	}

}
