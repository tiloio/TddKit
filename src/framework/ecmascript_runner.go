package main

import (
	"bytes"
	_ "embed"
	"log"
	"os/exec"

	"github.com/evanw/esbuild/pkg/api"
)

func ExecuteEcmascriptTests(file *string) *[]byte {

	log.Println("FILE", *file)

	result := api.Build(esBuildOptions(file))

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
		cmd = exec.Command("node", "--enable-source-maps", "--input-type=module", "-")
	} else {
		cmd = exec.Command("node", "--enable-source-maps", "-")
	}

	buffer := bytes.Buffer{}
	buffer.Write(result.OutputFiles[0].Contents)
	cmd.Stdin = &buffer
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

func esBuildOptions(file *string) api.BuildOptions {
	var options = api.BuildOptions{
		EntryPoints:   []string{*file},
		Write:         false,
		Outdir:        "out",
		Bundle:        true,
		Platform:      api.PlatformNode,
		Sourcemap:     api.SourceMapInline,
		Format:        api.FormatCommonJS,
		LegalComments: api.LegalCommentsNone,
	}

	if *EsModule {
		options.Format = api.FormatESModule
	}

	if jestLegacy := JestLegacyInjections(); jestLegacy != nil {
		options.Banner = map[string]string{
			"js": *jestLegacy.prefix,
		}
		options.Footer = map[string]string{
			"js": *jestLegacy.suffix,
		}
	}

	return options
}
