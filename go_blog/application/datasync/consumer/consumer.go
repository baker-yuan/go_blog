package consumer

import pb "github.com/baker-yuan/go-blog/protocol/datasync"

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
	}
}

// Send 将 TableChange 消息发送到通道
func Send(tableChange *pb.TableChange) {
	tableChangeChan <- tableChange
}
