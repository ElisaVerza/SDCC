package main

import (
	"strconv"
	"net/rpc"
	"fmt"
    "os"
)

func rpc_map(part string, cli *rpc.Client){
	var reply string
	fmt.Println(part)
	cli.Call("API.FirstTry", part, &reply)
	fmt.Println(reply)
}

func Divide(path string, cli[]*rpc.Client){
    var read, err = os.ReadFile(path)
	fmt.Println(err)
	for i:=0; i<len(cli); i++{
		go rpc_map(string(read[0:5]), cli[i])
	}
	for{}
}


func main() {
	var err error
	var num int
	num, _ = strconv.Atoi(os.Args[2])
	clients := make([] *rpc.Client, num)
	
	for i:=0 ; i < num; i++{
		var port = 4041+i
		clients[i], err = rpc.DialHTTP("tcp", "localhost:"+string(port)) /* create connection with server */
		if err != nil {
			fmt.Println("Error in connection , retring ... ") /*Retry connection every 3 seconds if server is not responding*/
		}
	}
	Divide(os.Args[1], clients)
}
