package ziface

type IServer interface {

	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Serve 开启业务服务方法
	Serve()
}
