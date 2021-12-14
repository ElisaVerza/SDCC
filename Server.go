package main

import (
	"strconv"
	"net/rpc"
	"bufio"
	"fmt"
	"log"
    "os"
)

func rpc_map(part string, cli *rpc.Client, c chan map[string]int){
	var reply map[string]int
	cli.Call("API.FirstTry", part, &reply)
	//fmt.Println(reply)
	c <- reply
}

func Divide(path string, cli[]*rpc.Client){
	var text []string
	mapResult1:= make(map[string]int)

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
//Creazione silce con tutte le map dei mapper
	for{
		mapResult := <-c
		for key , _/*value*/ :=  range mapResult {
			fmt.Println(key)
			mapResult1[key] = 1234//mapResult1[key].append(value)
		}
		fmt.Println(mapResult1)

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
