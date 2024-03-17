package go_mysql

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/datasync/config"
	"github.com/baker-yuan/go-blog/application/blog/datasync/consumer"
	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/datasync"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
)

// EventHandler 时间处理
type EventHandler struct {
	canal.DummyEventHandler
}

func (h *EventHandler) convertToString(value interface{}) (string, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case []byte:
		return string(v), nil
	case string:
		return v, nil
	case time.Time:
		return util.DateUtils.FormatDateTime(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	//case nil:
	//	return "NULL", nil
	default:
		// 对于复杂类型，如时间、结构体等，需要根据实际情况进行处理
		//return fmt.Sprintf("%v", v), nil
		return "", fmt.Errorf("unsupported value type: %T", v)
	}
}

// OnRow 当行事件（INSERT、UPDATE、DELETE）发生时调用。RowsEvent 包含了变更的行数据。
func (h *EventHandler) OnRow(e *canal.RowsEvent) error {
	tableChange := pb.TableChange{
		DbName: e.Table.Schema,
		TbName: e.Table.Name,
	}
	switch e.Action {
	case canal.UpdateAction:
		// 对于 update，Rows 包含成对的行数据：[旧值, 新值]
		before := e.Rows[0] // 旧值
		after := e.Rows[1]  // 新值
		beforeColumnMap := make(map[string]string)
		afterColumnMap := make(map[string]string)
		for i, col := range e.Table.Columns {
			beforeColumnMap[col.Name], _ = h.convertToString(before[i])
			afterColumnMap[col.Name], _ = h.convertToString(after[i])
		}
		tableChange.BeforeColumnMap = beforeColumnMap
		tableChange.AfterColumnMap = afterColumnMap
	case canal.InsertAction:
		columnMap := make(map[string]string)
		for i, col := range e.Table.Columns {
			columnMap[col.Name], _ = h.convertToString(e.Rows[0][i])
		}
		tableChange.AfterColumnMap = columnMap
	case canal.DeleteAction:
		columnMap := make(map[string]string)
		for i, col := range e.Table.Columns {
			columnMap[col.Name], _ = h.convertToString(e.Rows[0][i])
		}
		tableChange.BeforeColumnMap = columnMap
	}

	if len(tableChange.AfterColumnMap) != 0 {
		tableChange.ColumnMap = tableChange.AfterColumnMap
	} else {
		tableChange.ColumnMap = tableChange.BeforeColumnMap
	}
	// 使用 consumer 包的 Send 方法将解析出的数据发送到通道
	consumer.Send(&tableChange)
	return nil
}

// String 返回一个代表事件处理器状态的字符串，通常用于日志记录或调试。
func (h *EventHandler) String() string {
	return "EventHandler"
}

// getLatestPosition 获取最新的位置
func getLatestPosition(ctx context.Context, cfg *config.Config) (*mysql.Position, error) {
	// 连接到 MySQL 服务器
	conn, err := client.Connect(cfg.Mysql.Addr, cfg.Mysql.User, cfg.Mysql.Password, "")
	if err != nil {
		fmt.Println("failed to connect to MySQL:", err)
		return nil, err
	}
	defer conn.Close()

	// 执行 SHOW MASTER STATUS 命令
	r, err := conn.Execute("SHOW MASTER STATUS")
	if err != nil {
		fmt.Println("failed to execute SHOW MASTER STATUS:", err)
		return nil, err
	}

	// 获取最新的 binlog 文件名和位置
	binlogFile, _ := r.GetString(0, 0)
	binlogPos, _ := r.GetInt(0, 1)

	return &mysql.Position{Name: binlogFile, Pos: uint32(binlogPos)}, nil // 从最新的 binlog 开始，或者指定位置开始
}

func Init(ctx context.Context, cfg *config.Config) error {
	// 确定监听的binlog起点
	var (
		position *mysql.Position
		err      error
	)
	if len(cfg.BinLog.PositionName) != 0 && cfg.BinLog.PositionPos != 0 {
		position = &mysql.Position{Name: cfg.BinLog.PositionName, Pos: cfg.BinLog.PositionPos}
	} else {
		position, err = getLatestPosition(ctx, cfg)
		if err != nil {
			panic(err)
		}
	}
	// 监听配置
	canalCfg := canal.NewDefaultConfig()
	canalCfg.Addr = cfg.Mysql.Addr
	canalCfg.User = cfg.Mysql.User
	canalCfg.Password = cfg.Mysql.Password
	if len(cfg.Mysql.DbName) != 0 {
		canalCfg.Dump.TableDB = cfg.Mysql.DbName
	}
	if len(cfg.Mysql.TbName) != 0 {
		canalCfg.Dump.Tables = cfg.Mysql.TbName
	}
	// 启动监听
	// 此canal非阿里的canal，这里模仿了下阿里的canal，同名了
	c, err := canal.NewCanal(canalCfg)
	if err != nil {
		panic(err)
	}
	c.SetEventHandler(&EventHandler{})
	if err = c.RunFrom(*position); err != nil {
		panic(err)
	}
	return nil
}
