package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var version = flag.Bool("version", false, "print the version")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("protoc-gen-go-errors %v\n", release)
		return
	}
	var flags flag.FlagSet
	pkg := flags.String("pkg", "", "errors pkg")
	tml := flags.String("tml", "", "tml file path")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		errorsPackage = (*protogen.GoImportPath)(pkg)
		if *errorsPackage == "" {
			panic("need error pkg. e.g., --nilnoun-errors_opt=pkg/errors")
		}
		if err := InitErrorsTemplate(*tml); err != nil {
			panic(fmt.Sprintf("failed to load tml file, err:%s", err))
		}

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}
