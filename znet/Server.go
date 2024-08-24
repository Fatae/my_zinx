package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {

	// 服务器的名称
	Name string
	// tcp4 or other
	IPVersion string
	// 服务绑定的ip地址
	IP string
	// 服务绑定的端口
	Port int
}

func (s *Server) Start() {
	fmt.Println("Server start...")

	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))

		if err != nil {
			fmt.Println("resolve ip addr error:", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen error:", err)
			return
		}

		fmt.Println("start zinx server", s.Name, "success, listen on")

		for {
			conn, err := listener.AcceptTCP()

			if err != nil {
				fmt.Println("accept error:", err)
				continue
			}

			go func() {

				for {
					buf := make([]byte, 512)

					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read error:", err)
						continue
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back error:", err)
						continue
					}
				}

			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {

	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      7777,
	}
	return s
}
