package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {

	nodeFile, err := os.ReadFile("./index.js")
	check(err)

	oneGoRoutineEachCPu(nodeFile)
	// oneGoRoutineEachTest(nodeFile)
}

func oneGoRoutineEachTest(nodeFile []byte) {
	var channels = make(chan bool, 100)

	for i := 0; i < 100; i++ {
		go execute(nodeFile, channels)
	}
	for i := 0; i < 100; i++ {
		<-channels
	}
}

func oneGoRoutineEachCPu(nodeFile []byte) {
	var cpus = runtime.NumCPU() * 2
	var channels = make(chan bool, cpus)

	for i := 0; i < cpus; i++ {
		go execute(nodeFile, channels)
	}
	for i := 0; i < cpus; i++ {
		<-channels
	}

}

func execute(nodeFile []byte, channel chan bool) {
	var cmd *exec.Cmd
	cmd = exec.Command("node", "--enable-source-maps", "-")

	buffer := bytes.Buffer{}
	buffer.Write(nodeFile)
	cmd.Stdin = &buffer
	out, err := cmd.CombinedOutput()
	fmt.Printf("combined out:\n%s\n", string(out))
	check(err)
	channel <- true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
