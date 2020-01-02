package main

import "github.com/lihuicms-code-rep/goGameServer/znet"

func main() {
	s := znet.NewServer()
	s.Serve()
}

