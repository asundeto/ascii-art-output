package main

import (
	ascart "ascart/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	var param []string
	for i := 1; i < len(os.Args); i++ {
		param = append(param, os.Args[i])
	}
	output, content, style := ascart.Param(param)
	if ascart.CheckTxt(output) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if style == "" {
		style = "standard"
	}
	if ascart.CheckFormat(style) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if !ascart.CheckHash(style) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if output != "" {
		ascart.WriteToFile(output, ascart.ReadFile(content, style))
		fmt.Print(ascart.ReadFile(content, style))
	} else {
		fmt.Print(ascart.ReadFile(content, style))
	}
}
