# TCP Echo 服务器

这是一个简单的 TCP Echo 服务器，用 Go 语言实现。它可以接收任何 TCP 连接，并将接收到的数据原样返回给客户端。服务器会实时显示所有连接状态和数据传输情况。

## 功能特点

- 支持自定义监听端口
- 支持多客户端并发连接
- 详细的连接状态日志
- 数据收发统计
- 实时显示数据内容

## 运行要求

- Go 1.21 或更高版本

## 使用方法

### 编译和运行

1. 直接运行（使用默认端口 8080）：
```bash
go run main.go
```

2. 指定自定义端口运行：
```bash
go run main.go -port 9000
```

### 命令行参数

- `-port`: 指定服务器监听的端口号（默认：8080）

### 测试服务器

可以使用多种工具测试服务器：

1. 使用 netcat (nc)：
```bash
nc localhost 8080
```

2. 使用 telnet：
```bash
telnet localhost 8080
```

3. 使用其他 TCP 客户端工具

连接后，输入任何文本内容，服务器都会将其原样返回。

## 输出说明

服务器的输出使用以下格式：

- `[时间戳] Echo 服务器已启动，监听端口 xxxx...`
- `[时间戳] 新连接建立: <客户端地址>`
- `[时间戳] ◀ 收到来自 <客户端地址> 的数据 (xx 字节): <数据内容>`
- `[时间戳] ▶ 发送到 <客户端地址> 的数据 (xx 字节): <数据内容>`
- `[时间戳] 客户端断开连接 [<客户端地址>], 接收: xx 字节, 发送: xx 字节`
- `[时间戳] 连接关闭: <客户端地址>`

### 符号说明
- ◀：表示接收到的数据
- ▶：表示发送的数据

## 示例输出

```
[2024-03-19 15:04:05] Echo 服务器已启动，监听端口 8080...
[2024-03-19 15:04:10] 新连接建立: 127.0.0.1:52431
[2024-03-19 15:04:12] ◀ 收到来自 127.0.0.1:52431 的数据 (5 字节): hello
[2024-03-19 15:04:12] ▶ 发送到 127.0.0.1:52431 的数据 (5 字节): hello
[2024-03-19 15:04:15] 客户端断开连接 [127.0.0.1:52431], 接收: 5 字节, 发送: 5 字节
[2024-03-19 15:04:15] 连接关闭: 127.0.0.1:52431
```

## 注意事项

1. 服务器会一直运行，直到被手动终止（Ctrl+C）
2. 每个客户端连接都在独立的 goroutine 中处理，支持并发连接
3. 默认的接收缓冲区大小为 1024 字节 