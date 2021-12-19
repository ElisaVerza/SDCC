package main

import (
	"strconv"
	"net/rpc"
	"bufio"
	"fmt"
	"log"
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
//Funzione eseguita dai thread figli per passare ai vari mapper la porzione di file da analizzare
//Il thread resta in attesa della risposta del mapper e successivamente la comunica attraverso un 
//channel al main thread.

func rpc_map(part string, cli *rpc.Client, c chan map[string]int){
	var reply map[string]int
	cli.Call("API.Mapper", part, &reply)
	c <- reply
}

//Funzione che gestisce il master, si occupa di leggere e suddividere il lavoro tra i worker sia
//della fase di map che della fase di reduce. Effettua tra le due fasi shuffle and sort.

func Divide(path string, cli[]*rpc.Client){	
	var text []string
	var reply WordCounted
	sorter:= make(map[string][]int)
	
//Creazione file dove verrà salvato risultato finale
	f, err := os.Create("finalGrep.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	nodes := len(cli) 
	c := make(chan map[string]int)
	
//Apertura file.
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open")

	}
	
//Divisione file in righe.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	chunk := len(text)
	file.Close()
//Assegnazione righe ai mapper.
	j:=0
	for i:=0; i<nodes; i++{
		go rpc_map(text[j], cli[i], c)
		if i==(nodes-1) && j<chunk{
			i = 0
		}
		if j == (chunk-1){
			break
		}
		j++

	}
	
//Creazione map con key la parola e value una slice con tutte la occorrenze della parola trovata dai mapper (non sommate)
	for i:=0; i<chunk; i++{
		mapResult := <-c
		for key , value :=  range mapResult {
			sorter[key] = append(sorter[key], value)
		}
	}

//Distribuzione lavoro tra i reducer e scrittura su file del risultato finale
	count:=0
	for k , v := range sorter{
		nodeWord := Word{Key:k, Value:v}
		if count==(nodes-1){
			count = 0
		}
		rpcCli := cli[count]
		rpcCli.Call("API.Reducer", nodeWord, &reply)
		count++

		_, err = f.WriteString(reply.KeyFinal+" ")
		_, err = f.WriteString(strconv.Itoa(reply.ValueCounted))
		_, err = f.WriteString("\n")


		if err != nil {
			log.Fatal(err)
			}	
	}
	fmt.Println("Grep eseguito")

}

//Nel main viene effettuata solo la connessione con i worker per le chiamate rpc
func main() {
	var err error
	var num int
	var port int
	
	num, _ = strconv.Atoi(os.Args[2])
	clients := make([] *rpc.Client, num)
	
	for i:=0 ; i < num; i++{
		port = 4041+i
		clients[i], err = rpc.DialHTTP("tcp", "localhost:"+strconv.Itoa(port))
		if err != nil {
			fmt.Println("Errore di connessione , retring ... ")
		}
	}
	Divide(os.Args[1], clients)
}
