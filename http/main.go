package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 定义处理函数
	http.HandleFunc("/", handleEcho)
	// 定义命令行参数
	port := flag.Int("port", 8080, "指定监听的端口号")
	flag.Parse()
	fmt.Printf("[%s] HTTP Echo 服务器已启动，监听端口 %d...\n",
		time.Now().Format("2006-01-02 15:04:05"), *port)

	// 启动服务器
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.RemoteAddr
	fmt.Printf("[%s] 新请求: %s %s\n",
		time.Now().Format("2006-01-02 15:04:05"), r.Method, remoteAddr)

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("[%s] 读取请求体错误 [%s]: %v\n",
			time.Now().Format("2006-01-02 15:04:05"), remoteAddr, err)
		http.Error(w, "读取请求失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	receivedBytes := len(body)
	fmt.Printf("[%s] ◀ 收到来自 %s 的数据 (%d 字节): %s\n",
		time.Now().Format("2006-01-02 15:04:05"), remoteAddr, receivedBytes, string(body))

	// 回显数据
	w.Header().Set("Content-Type", "text/plain")
	// 模拟耗时请求
	// 获取参数 delay
	delay := r.URL.Query().Get("delay")
	if delay != "" {
		delayFloat, err := strconv.ParseFloat(delay, 64)
		if err != nil {
			fmt.Printf("[%s] 无效的延迟参数 [%s]: %v\n",
				time.Now().Format("2006-01-02 15:04:05"), remoteAddr, err)
			delayFloat = 0
		}
		fmt.Printf("[%s] 延迟 %f 秒\n", time.Now().Format("2006-01-02 15:04:05"), delayFloat)
		time.Sleep(time.Duration(delayFloat * float64(time.Second)))
	}

	written, err := w.Write(body)
	if err != nil {
		fmt.Printf("[%s] 发送错误 [%s]: %v\n",
			time.Now().Format("2006-01-02 15:04:05"), remoteAddr, err)
		return
	}

	fmt.Printf("[%s] ▶ 发送到 %s 的数据 (%d 字节): %s\n",
		time.Now().Format("2006-01-02 15:04:05"), remoteAddr, written, string(body))

	fmt.Printf("[%s] 请求完成 [%s], 接收: %d 字节, 发送: %d 字节\n",
		time.Now().Format("2006-01-02 15:04:05"), remoteAddr, receivedBytes, written)
}
