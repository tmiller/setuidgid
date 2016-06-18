package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"syscall"
)

var usage string = "setuidgid: usage: setuidgid username program [arg...]"

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(111)
	}
}

func main() {

	if len(os.Args) <= 2 {
		fmt.Print(usage)
		os.Exit(100)
	}

	username := os.Args[1]
	program := os.Args[2]
	pargv := os.Args[2:]

	user, err := user.Lookup(username)
	checkError(err)

	uid, err := strconv.Atoi(user.Uid)
	checkError(err)
	gid, err := strconv.Atoi(user.Gid)
	checkError(err)

	err = syscall.Setgid(gid)
	checkError(err)
	err = syscall.Setuid(uid)
	checkError(err)

	if path.IsAbs(program) {
		err := syscall.Exec(program, pargv, os.Environ())
		checkError(err)
	}

	for _, p := range strings.Split(os.Getenv("PATH"), ":") {
		absPath := path.Join(p, program)
		err = syscall.Exec(absPath, pargv, os.Environ())
	}

	checkError(err)
}
