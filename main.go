package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 定义命令行参数
	port := flag.Int("port", 8080, "指定监听的端口号")
	flag.Parse()

	// 监听指定端口
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
	defer listener.Close()

	fmt.Printf("[%s] Echo 服务器已启动，监听端口 %d...\n", 
		time.Now().Format("2006-01-02 15:04:05"), *port)

	for {
		// 接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[%s] 接受连接失败: %v\n", 
				time.Now().Format("2006-01-02 15:04:05"), err)
			continue
		}

		// 为每个连接启动一个新的 goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("[%s] 新连接建立: %s\n", 
		time.Now().Format("2006-01-02 15:04:05"), remoteAddr)

	defer func() {
		conn.Close()
		fmt.Printf("[%s] 连接关闭: %s\n", 
			time.Now().Format("2006-01-02 15:04:05"), remoteAddr)
	}()

	// 创建一个缓冲区
	buffer := make([]byte, 1024)
	
	// 记录连接的数据统计
	var totalBytesReceived, totalBytesSent int64

	for {
		// 读取数据
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("[%s] 读取错误 [%s]: %v\n", 
					time.Now().Format("2006-01-02 15:04:05"), remoteAddr, err)
			} else {
				fmt.Printf("[%s] 客户端断开连接 [%s], 接收: %d 字节, 发送: %d 字节\n", 
					time.Now().Format("2006-01-02 15:04:05"), remoteAddr, 
					totalBytesReceived, totalBytesSent)
			}
			return
		}

		totalBytesReceived += int64(n)

		// 打印接收到的数据到标准输出
		fmt.Printf("[%s] ◀ 收到来自 %s 的数据 (%d 字节): %s", 
			time.Now().Format("2006-01-02 15:04:05"), remoteAddr, n, string(buffer[:n]))

		// 将数据回显给客户端
		written, err := conn.Write(buffer[:n])
		if err != nil {
			fmt.Printf("[%s] 发送错误 [%s]: %v\n", 
				time.Now().Format("2006-01-02 15:04:05"), remoteAddr, err)
			return
		}
		
		totalBytesSent += int64(written)
		
		// 打印发送的数据
		fmt.Printf("[%s] ▶ 发送到 %s 的数据 (%d 字节): %s", 
			time.Now().Format("2006-01-02 15:04:05"), remoteAddr, written, string(buffer[:written]))
	}
} 