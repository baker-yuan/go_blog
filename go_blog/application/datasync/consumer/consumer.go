package consumer

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/datasync"
	"trpc.group/trpc-go/trpc-go/client"
)

// tableChangeChan 用于传递TableChange消息的通道
var tableChangeChan chan *pb.TableChange

// Init 初始化通道和处理逻辑
//
// bufferSize 通道的缓冲大小
func Init(bufferSize uint32) {
	tableChangeChan = make(chan *pb.TableChange, bufferSize)
	go processTableChanges()
}

// processTableChanges 从通道读取数据并发送到消息队列
func processTableChanges() {

	for tableChange := range tableChangeChan {
		// 发送 tableChange 到消息队列
		// ... 消息队列发送逻辑 ...
		tableChange = tableChange

		// 创建客户端
		opts := []client.Option{
			client.WithMaxWindowSize(1 * 1024 * 1024),
		}
		proxy := pb.NewDataSyncApiClientProxy(opts...)
		streamClient, err := proxy.DataChange(context.Background())
		if err != nil {
			continue
		}
		// 循环发送
		for i := 0; i < 1; i++ {
			if err := streamClient.Send(&pb.TableChange{}); err != nil {
				continue
			}
		}
		// 关闭流
		_, err = streamClient.CloseAndRecv()
		if err != nil {
			continue
		}

	}
}

// Send 将 TableChange 消息发送到通道
func Send(tableChange *pb.TableChange) {
	tableChangeChan <- tableChange
}
