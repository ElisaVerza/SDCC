package main

import (
	"strconv"
	"net/rpc"
	"bufio"
	"fmt"
	"log"
    "os"
)

func rpc_map(part string, cli *rpc.Client){
	var reply map[string]int
	cli.Call("API.FirstTry", part, &reply)
	fmt.Println(reply)
}

func Divide(path string, cli[]*rpc.Client){
	
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	j:=0
	for i:=0; i<len(cli); i++{
		go rpc_map(text[j], cli[i])
		if i==(len(cli)-1) && j!=len(text){
			i = 0
		}else if j == (len(text)-1){
			break
		}
		j++
	}
	for{}
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
