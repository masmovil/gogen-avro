package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/actgardner/gogen-avro/v10/generator"
	"github.com/actgardner/gogen-avro/v10/generator/flat"
	"github.com/actgardner/gogen-avro/v10/parser"
	"github.com/actgardner/gogen-avro/v10/resolver"
)

const fileComment = "// Code generated by github.com/masmovil/gogen-avro/v10. DO NOT EDIT."

func main() {
	cfg := parseCmdLine()

	var err error
	pkg := generator.NewPackage(cfg.packageName, codegenComment(cfg))
	namespace := parser.NewNamespace(cfg.shortUnions)
	gen := flat.NewFlatPackageGenerator(pkg, cfg.containers)

	switch cfg.namespacedNames {
	case nsShort:
		generator.SetNamer(generator.NewNamespaceNamer(true))
	case nsFull:
		generator.SetNamer(generator.NewNamespaceNamer(false))
	}

	for _, fileName := range cfg.files {
		schema, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %q - %v\n", fileName, err)
			os.Exit(2)
		}

		_, err = namespace.TypeForSchema(schema)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding schema for file %q - %v\n", fileName, err)
			os.Exit(3)
		}
	}

	for _, def := range namespace.Roots {
		if err := resolver.ResolveDefinition(def, namespace.Definitions); err != nil {
			fmt.Fprintf(os.Stderr, "Error resolving definition for type %q - %v\n", def.Name(), err)
			os.Exit(4)
		}
	}

	for _, def := range namespace.Roots {
		err = gen.Add(def)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating code for schema - %v\n", err)
			os.Exit(5)
		}
	}

	err = pkg.WriteFiles(cfg.targetDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing source files to directory %q - %v\n", cfg.targetDir, err)
		os.Exit(6)
	}
}

// codegenComment generates a comment informing readers they are looking at
// generated code. If source avro files are provided they're included in the comment
func codegenComment(c config) string {
	if !c.sourcesComment {
		return fileComment
	}
	sourcesComment := `%s
/*
 * %s
 */`
	var sourceBlock []string
	if len(c.files) == 1 {
		sourceBlock = append(sourceBlock, "SOURCE:")
	} else {
		sourceBlock = append(sourceBlock, "SOURCES:")
	}

	for _, source := range c.files {
		_, fName := filepath.Split(source)
		sourceBlock = append(sourceBlock, fmt.Sprintf(" *     %s", fName))
	}
	return fmt.Sprintf(sourcesComment, fileComment, strings.Join(sourceBlock, "\n"))

}
