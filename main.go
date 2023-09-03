package main

import (
	"os"
)

func main() {
	output, err := ExtractContent(os.Stdin)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.WriteString(output)
	if err != nil {
		panic(err)
	}
}
