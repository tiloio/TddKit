package main

import (
	_ "embed"
	"log"
	"path/filepath"
	"os/exec"
	// "bytes"
	"github.com/evanw/esbuild/pkg/api"
    "os"
)

//go:embed adapters/node/index.js
var nodeRequire2 string

func ExecuteEcmascriptTestsFs(file *string) *[]byte {

	log.Println("FILE", *file)

	format := api.FormatCommonJS

	if *EsModule {
		format = api.FormatESModule
	}

	tempDir, err := os.MkdirTemp("", "test-framework")
	defer os.RemoveAll(tempDir)

	const testFileName  = "test.js"
	var testFilePath = filepath.Join(tempDir, testFileName)

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{*file},
		Write:       true,
		Outfile: 	testFilePath,
		Bundle:      true,
		Platform: 	 api.PlatformNode,
		Sourcemap:   api.SourceMapInline,
		Format: 	 format,
		Banner: map[string]string{
			"js":  nodeRequire2,
		  },
	  })
	
	if len(result.Errors) > 0 {
		log.Println("Could not parse files via esbuild, got erros:")
		log.Fatal(result.Errors)
	}
	if len(result.OutputFiles) != 1 {
		log.Println("Could not parse files via esbuild, got zero or more than one OutputFile:")
		for _, out := range result.OutputFiles {
			log.Printf("%v\n%v\n", out.Path, string(out.Contents))
		}
		log.Fatal("")
	}

	var cmd *exec.Cmd
	if *EsModule {
		cmd = exec.Command("node", "--enable-source-maps", "--input-type=module", testFilePath)
	} else {
		cmd = exec.Command("node", "--enable-source-maps", testFilePath)
	}

	//cmd.Dir = filepath.Dir(*file)

	//buffer := bytes.Buffer{}
	//buffer.Write(result.OutputFiles[0].Contents)
	//cmd.Stdin = &buffer
	stdoutAndStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Command finished with error:", err)
		log.Println("Executed file:")
		log.Println(string(result.OutputFiles[0].Contents))
		log.Println("Command output")
		log.Fatal(string(stdoutAndStderr))
	}

	return &stdoutAndStderr
}
