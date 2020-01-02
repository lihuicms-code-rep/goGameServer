package ziface

//服务模块抽象层
type IServer interface {
	Start() //服务启动
	Stop()  //服务关闭
	Serve() //服务运行
}
