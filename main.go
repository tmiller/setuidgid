package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	//path := os.Getenv("PATH")

	user, err := user.Lookup(os.Args[1])
	checkError(err)

	uid, err := strconv.Atoi(user.Uid)
	checkError(err)

	gid, err := strconv.Atoi(user.Gid)
	checkError(err)

	err = syscall.Setuid(uid)
	checkError(err)

	err = syscall.Setgid(gid)
	checkError(err)

	if strings.HasPrefix(os.Args[2], "/") {
		err := syscall.Exec(os.Args[2], os.Args[2:], os.Environ())
		checkError(err)
	}
}
