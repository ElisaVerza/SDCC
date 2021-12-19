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

type API int


func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

//Funzione che esegue la fase di reduce, prende in input una struct che ha come key un parola
//e come value una lista con tutte le occorrenze da sommare della parola.
func (a *API) Reducer(w Word, reply *WordCounted) error {

	var finalStruct WordCounted 
	finalStruct.ValueCounted = 0
	finalStruct = WordCounted{KeyFinal:w.Key}
	
//Somma tutti i valori nella lista, restituisce una struct con key la parola e value la somma finale
	for i:=0; i<len(w.Value); i++{
		finalStruct.ValueCounted += w.Value[i]
	}
    *reply = finalStruct 
	return nil


}

//Funzione che esegue la fase di map, prende come input la stringa da analizzare
func (a *API) Mapper(i string, reply *map[string]int) error {
	
//Formattazione del testo
    reg, err := regexp.Compile(`[^\d\p{Latin}]`)
	if err != nil {
        log.Fatal(err)
    }
    str := reg.ReplaceAllString(i, " ")
	str = strings.ToLower(str)

//Creazione map con key la parola e value il numero di volte che che compare nel proprio chunk
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

	if err != nil {
		log.Fatal("Errore nella registrazione mapper", err)
	}

//Consumer in attesa di ricevere chiamate rpc
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", string(os.Args[1]))
	if err != nil {
		log.Fatal("Errore nella registrazione listener", err)
	}

	log.Printf("Serving rpc sulla porta %s", string(os.Args[1]))
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Errore in serving : ", err)
	}

}