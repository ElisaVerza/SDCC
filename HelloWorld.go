// hello.go

package main

import (
	"strings"
    "fmt"
    "os"
)

var path = "testo.txt"

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

func main() {
    fmt.Println("Opening a file ")
    var file, err = os.ReadFile(path)
    if isError(err) {
        return
    }
	lowerString := strings.ToLower(string(file))
    arrayString := strings.Fields(lowerString)
	dict:= make(map[string]int)
    for _ , num :=  range arrayString {
        dict[num] = dict[num]+1
    }
    fmt.Println(dict)
}