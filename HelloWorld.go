// hello.go

package main

import (
	"strings"
    "fmt"
	"log"
    "os"
	"regexp"
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
	
	reg, err := regexp.Compile(`[^\d\p{Latin}]`)
    if err != nil {
        log.Fatal(err)
    }
    str := reg.ReplaceAllString(string(file), " ")
	str = strings.ToLower(str)
	
    arrayString := strings.Fields(str)
	dict:= make(map[string]int)
    for _ , num :=  range arrayString {
        dict[num] = dict[num]+1
    }
    fmt.Println(dict)
}