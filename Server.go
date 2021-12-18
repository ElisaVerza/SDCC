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

func rpc_map(part string, cli *rpc.Client, c chan map[string]int){
	var reply map[string]int
	cli.Call("API.Mapper", part, &reply)
	//fmt.Println(reply)
	c <- reply
}

func Divide(path string, cli[]*rpc.Client){	
	var text []string
	var reply WordCounted
	sorter:= make(map[string][]int)

	nodes := len(cli) 
	c := make(chan map[string]int)
//Apertura file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open")

	}
//Divisione file in righe
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	chunk := len(text)
	file.Close()
//Assegnazione righe ai mapper
	j:=0
	for i:=0; i<nodes; i++{
		go rpc_map(text[j], cli[i], c)
		if i==(nodes-1) && j!=chunk{
			i = 0
		}else if j == (chunk-1){
			break
		}
		j++
	}
//Map con key la parola e value una slice con tutte la occorrenze della parola trovata dai mapper (non sommate)
	for i:=0; i<chunk; i++{
		mapResult := <-c
		for key , value :=  range mapResult {
			sorter[key] = append(sorter[key], value)
		}
	}
	count:=0
	for k , v := range sorter{
		nodeWord := Word{Key:k, Value:v}
		if count==(nodes-1){
			count = 0
		}
		rpcCli := cli[count]
		rpcCli.Call("API.Reducer", nodeWord, &reply)
		count++

		fmt.Println(reply.ValueCounted)
	}

}


func main() {
	var err error
	var num int
	var port int
	
	num, _ = strconv.Atoi(os.Args[2])
	clients := make([] *rpc.Client, num)
	
	for i:=0 ; i < num; i++{
		port = 4041+i
		clients[i], err = rpc.DialHTTP("tcp", "localhost:"+strconv.Itoa(port)) /* create connection with server */
		if err != nil {
			fmt.Println("Error in connection , retring ... ") /*Retry connection every 3 seconds if server is not responding*/
		}
	}
	Divide(os.Args[1], clients)
}
