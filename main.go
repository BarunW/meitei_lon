package main

import (
	"fmt"
	"meitei_lon/repl"
	"os"
)

func main(){
    fmt.Println("Happy New Year 2025")
    repl.Start(os.Stdin, os.Stdout)
}


