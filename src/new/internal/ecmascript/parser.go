package ecmascript

import (
	"flag"
	"test-framework/internal/logger"
	"test-framework/internal/parser"

	"github.com/evanw/esbuild/pkg/api"
)

var esModule = flag.Bool("esm", false, "Is code es module?")

func ParseFiles(files *[]string) *[]parser.File {
	var filesLength = len(*files)
	parsedFileChannel := make(chan parser.File, filesLength)

	for _, file := range *files {
		go parseFile(file, parsedFileChannel)
	}

	var parsedFiles = make([]parser.File, filesLength)

	for i := 0; i < filesLength; i++ {
		parsedFiles[i] = <-parsedFileChannel
	}

	return parsedFiles
}

func parseFile(file string, fileContent chan parser.File) {
	result := api.Build(esBuildOptions(file))

	if len(result.Errors) > 0 {
		logger.Fatal("Could not parse files via esbuild, got erros: \n", result.Errors)
	}
	if len(result.OutputFiles) != 1 {

		var filesMessage = ""
		for _, out := range result.OutputFiles {
			filesMessage += out.Path + "\n"
		}
		logger.Fatal("Could not parse files via esbuild, got zero or more than one OutputFile: \n", filesMessage)
	}

	content := result.OutputFiles[0].Contents

	fileContent <- parser.File{
		Name:    file,
		content: content,
	}
}

func esBuildOptions(file string) api.BuildOptions {
	var options = api.BuildOptions{
		EntryPoints:   []string{file},
		Write:         false,
		Outdir:        "out",
		Bundle:        true,
		Platform:      api.PlatformNode,
		Sourcemap:     api.SourceMapInline,
		Format:        api.FormatCommonJS,
		LegalComments: api.LegalCommentsNone,
	}

	if *esModule {
		options.Format = api.FormatESModule
	}

	if jestLegacy := jestLegacyInjections(); jestLegacy != nil {
		options.Banner = map[string]string{
			"js": *jestLegacy.prefix,
		}
		options.Footer = map[string]string{
			"js": *jestLegacy.suffix,
		}
	}

	return options
}
