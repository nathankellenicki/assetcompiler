package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	files := os.Args[1:]
	output, err := os.Create("./assets.go")
	if err != nil {
		panic(err.Error())
	}

	output.Write([]byte("package main\n\n"))
	output.Write([]byte("var AssetData = map[string][]byte{\n"))

	for _, path := range files {
		filename := filepath.Base(path)
		fmt.Printf("%s...", filename)

		if _, err := os.Stat(path); err != nil {
			panic(err.Error())
		} else if content, err := ioutil.ReadFile(path); err == nil {

			var dataSlice []string

			output.Write([]byte(fmt.Sprintf("    \"%s\": {", filename)))

			for _, b := range content {
				byteString := fmt.Sprintf("0x%x", b)
				dataSlice = append(dataSlice, byteString)
			}

			dataString := strings.Join(dataSlice, ", ")

			output.WriteString(dataString)
			output.Write([]byte("},\n"))
			fmt.Print("done\n")

		} else {
			panic(err.Error())
		}

	}

	output.Write([]byte("}\n"))

}
