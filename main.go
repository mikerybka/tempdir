package main

import (
	"fmt"
	"os"
)

func main() {
	tempdir := os.TempDir()
	fmt.Println(tempdir)
}
