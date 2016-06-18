package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"syscall"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(111)
	}
}

func main() {

	username := os.Args[1]
	program := os.Args[2]

	user, err := user.Lookup(username)
	checkError(err)

	uid, err := strconv.Atoi(user.Uid)
	checkError(err)

	gid, err := strconv.Atoi(user.Gid)
	checkError(err)

	err = syscall.Setuid(uid)
	checkError(err)

	err = syscall.Setgid(gid)
	checkError(err)

	if path.IsAbs(program) {
		err := syscall.Exec(program, os.Args[2:], os.Environ())
		checkError(err)
	}
}
