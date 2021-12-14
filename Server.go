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
	fmt.Println(reply)
	c <- reply
}

func Divide(path string, cli[]*rpc.Client){
	var text []string
	nodes := len(cli) 
	c := make(chan map[string]int)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	chunk := len(text)
	file.Close()
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
	for {
		x, y, z := <-c, <-c, <-c
		fmt.Println("main", x)
		fmt.Println("main", y)
		fmt.Println("main", z)
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
