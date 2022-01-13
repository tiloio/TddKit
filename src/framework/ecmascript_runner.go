package main

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"os/exec"
)

func ExecuteEcmascriptTests(testCode *[]byte, environmentVariables *[]string) *[]byte {

	var cmd *exec.Cmd
	if *esModule {
		cmd = exec.Command("node", "--enable-source-maps", "--input-type=module", "-")
	} else {
		cmd = exec.Command("node", "--enable-source-maps", "-")
	}
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, *environmentVariables...)

	buffer := bytes.Buffer{}
	buffer.Write(*testCode)
	cmd.Stdin = &buffer
	stdoutAndStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Command finished with error:", err)
		log.Println("Executed file:")
		log.Println(string(*testCode))
		log.Println("Command output")
		log.Fatal(string(stdoutAndStderr))
	}

	return &stdoutAndStderr
}
