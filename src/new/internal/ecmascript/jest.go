package ecmascript

import (
	_ "embed"
	"flag"
)

var jestLegacyFlag = flag.Bool("jest", false, "Enables running tests written in the jest format.")

//go:embed node/dist/jest-legacy-prefix.js
var prefixInjections string
var suffixInjections string

type JestLegacyInjection struct {
	prefix *string
	suffix *string
}

var injections = JestLegacyInjection{
	prefix: &prefixInjections,
	suffix: &suffixInjections,
}

func jestLegacyInjections() *JestLegacyInjection {
	if *jestLegacyFlag {
		return &injections
	}

	return nil
}
