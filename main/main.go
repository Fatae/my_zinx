package main

import "zinx/znet"

func main() {
	s := znet.NewServer("zinxV0.1")

	s.Serve()
}
