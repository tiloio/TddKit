package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	nodeFile, err := os.ReadFile("./index.js")
	check(err)

	var cmd *exec.Cmd
	cmd = exec.Command("node", "--enable-source-maps", "-")

	buffer := bytes.Buffer{}
	buffer.Write(nodeFile)
	cmd.Stdin = &buffer
	out, err := cmd.CombinedOutput()
	fmt.Printf("combined out:\n%s\n", string(out))
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
