// hello.go

package main

import (
    "fmt"
	"log"
    "os"
	"strconv"
)

type wordCounted struct {
    keyFinal string
    valueCounted int
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

func main() {
	var finalStruct wordCounted 
	finalStruct.valueCounted = 0
	finalStruct.keyFinal = "Ciao"
	f, err := os.Create("c1.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
	for i:=0; i<10; i++{
		_, err = f.WriteString(finalStruct.keyFinal)
		_, err = f.WriteString(strconv.Itoa(finalStruct.valueCounted))

		if err != nil {
			log.Fatal(err)
		}
		finalStruct.valueCounted += i
	}
	return
}