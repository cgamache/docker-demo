package main

import (
    "fmt"
    "github.com/stretchr/piglatin"
    "io/ioutil"
    "os"
)

func main() {
    bytes, _ := ioutil.ReadAll(os.Stdin)
    pl := piglatin.Translate(string(bytes))
    fmt.Printf("%s", pl)
}