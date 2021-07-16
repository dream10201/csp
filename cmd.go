package main

import (
	"os/exec"
)

const listAll = "pdbedit -L"

func check(user string) {
	ec := exec.Command(listAll)
	ec.String()
}
