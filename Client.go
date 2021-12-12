package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"fmt"
    "os"
)

/*API needed for rpc calls*/
type API int


func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

func (a *API) FirstTry(i string, reply *string) error {
    var str string
	str = "Ciao"
	i = i+str
	fmt.Println("Ciau")
    *reply = string(i)
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("Errore di connessione : ", err)
	}

	if err != nil {
		log.Fatal("Errore nella registrazione delle API", err)
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