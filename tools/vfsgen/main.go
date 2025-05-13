// generate.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

func main() {
	fs := http.Dir("./web/dist")

	if err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "./internal/delivery/grpc/handler/ui_vfs.go",
		PackageName:  "handler",
		VariableName: "UIStaticFiles",
	}); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
