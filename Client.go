package main

import (
	"regexp"
	"strings"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"fmt"
    "os"
)

type Word struct {
    Key string
    Value []int
}

type WordCounted struct {
    KeyFinal string
    ValueCounted int
}
var path string

/*API needed for rpc calls*/
type API int
type Arith int


func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

func (a *API) Reducer(w Word, reply *WordCounted) error {
	var finalStruct WordCounted 
	finalStruct.ValueCounted = 0
	finalStruct = WordCounted{KeyFinal:w.Key}
	/*f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()*/
	for i:=0; i<len(w.Value); i++{
		/*_, err = f.WriteString(finalStruct.keyFinal)
		_, err = f.WriteString(strconv.Itoa(finalStruct.valueCounted))

		if err != nil {
			log.Fatal(err)
		}*/
		finalStruct.ValueCounted += w.Value[i]
	}
    *reply = finalStruct 
	return nil


}


func (a *API) Mapper(i string, reply *map[string]int) error {

    reg, err := regexp.Compile(`[^\d\p{Latin}]`)
	if err != nil {
        log.Fatal(err)
    }
    str := reg.ReplaceAllString(i, " ")
	str = strings.ToLower(str)
	
    arrayString := strings.Fields(str)
	dict:= make(map[string]int)
    for _ , num :=  range arrayString {
        dict[num] = dict[num]+1
    }
    *reply = dict
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	path = os.Args[2]

	if err != nil {
		log.Fatal("Errore nella registrazione mapper", err)
	}

	/*Consumer is listening for calls*/
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", string(os.Args[1]))
	if err != nil {
		log.Fatal("errore nella registrazione del server", err)
	}

	log.Printf("serving rpc sulla porta %s", string(os.Args[1]))
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Errore in serving : ", err)
	}

}