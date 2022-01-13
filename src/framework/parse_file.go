package main

import (
	_ "embed"
	"log"

	"github.com/evanw/esbuild/pkg/api"
)

type ParsedFile struct {
	name    string
	content []byte
}

func ParseFileAsync(file string, isEsModule *bool, fileContent chan ParsedFile) {
	var content = parseFile(file, isEsModule)
	fileContent <- ParsedFile{
		name:    file,
		content: content,
	}
}

func parseFile(file string, isEsModule *bool) []byte {
	result := api.Build(esBuildOptions(file, isEsModule))

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

	return result.OutputFiles[0].Contents
}

func esBuildOptions(file string, isEsModule *bool) api.BuildOptions {
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

	if *isEsModule {
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
