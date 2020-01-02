package znet

import (
	"fmt"
	"net"
)

//服务模块实现
type Server struct {
	Name       string //服务名称
	TCPVersion string //网络版本
	TCPAddr    string //服务绑定地址
	TCPPort    int    //服务绑定端口
}

//创建服务器句柄
func NewServer() *Server {
	return &Server{
		Name:       "myGameServer",
		TCPVersion: "tcp4",
		TCPAddr:    "0.0.0.0",
		TCPPort:    8888,
	}
}

//服务启动
func (s *Server) Start() {
	//1.获取TCPAddr
	addr, err := net.ResolveTCPAddr(s.TCPVersion, fmt.Sprintf("%s:%d", s.TCPAddr, s.TCPPort))
	if err != nil {
		fmt.Println("resolve tcp addr error ", err)
		return
	}

	//2.监听地址
	listener, err := net.ListenTCP(s.TCPVersion, addr)
	if err != nil {
		fmt.Println("listen tcp error ", err)
		return
	}

	fmt.Printf("[Start] %s listening at %s:%d success\n", s.Name, s.TCPAddr, s.TCPPort)

	//3.等待客户端连接过来并处理相应业务
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept client connect error ", err)
			continue
		}

		go s.connHandler(conn)
	}
}

//服务关闭
func (s *Server) Stop() {
	fmt.Printf("[Stop] %s stop", s.Name)
}

//服务运行
func (s *Server) Serve() {
	s.Start()
}

//处理连接过来后的具体业务,目前版本只将客户端发送的数据进行回显
func (s *Server) connHandler(conn net.Conn) {
	if conn == nil {
		return
	}

	//不断从客户端读取数据
	//这里是阻塞的,假如不用go connHandler,那么第二个客户端一直都没有机会连接进来
	for {
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read data from conn error ", err)
			return
		}

		fmt.Println("receive data from client:", string(buf[:cnt]))

		//回写到该连接对端
		if _, err := conn.Write(buf[:cnt]); err != nil {
			fmt.Println("write buf to conn error ", err)
			return
		}
	}
}