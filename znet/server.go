package znet

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

}

func (s *Server) stop() {

}

func (s *Server) Server() {

}
