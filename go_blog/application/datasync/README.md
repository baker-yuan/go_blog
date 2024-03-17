


# 数据同步服务

https://github.com/alibaba/canal
https://github.com/go-mysql-org/go-mysql
https://github.com/zendesk/maxwell
https://github.com/go-mysql-org/go-mysql-elasticsearch
https://github.com/github/gh-ost


## 配置
```sql
create database if not exists baker_sys_canal default char set utf8mb4;
use baker_sys_canal;

CREATE TABLE `baker_canal_config`
(
`id`             int(10) unsigned NOT NULL AUTO_INCREMENT,
`db_name`        varchar(255)     NOT NULL DEFAULT '' COMMENT '库名',
`tb_name`        varchar(255)     NOT NULL DEFAULT '' COMMENT '表名',
`monitor_insert` bit(1)           NOT NULL DEFAULT b'0' COMMENT '是否监听插入',
`monitor_update` bit(1)           NOT NULL DEFAULT b'0' COMMENT '是否监听修改',
`monitor_delete` bit(1)           NOT NULL DEFAULT b'0' COMMENT '是否监听删除',
`create_time`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time`    timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `uk_db_name_tb_name` (`db_name`, `tb_name`)
) ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARSET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci COMMENT ='canal监听配置|baker.yuan|2022-02-19';


```

```sql
insert into baker_canal_config(id, db_name, tb_name, monitor_insert, monitor_update, monitor_delete)
value (1, 'baker_sys_canal', 'baker_canal_config', true, true, true);

insert into baker_canal_config(id, db_name, tb_name, monitor_insert, monitor_update, monitor_delete)
value (2, 'baker_blog_blog', 'blog_article', true, true, true);
```


## 基于 binlog 监听

基于 binlog 监听的数据同步工具专注于解析数据库的二进制日志（binlog），这是数据库用来记录所有更改（如插入、更新和删除）的日志文件。以下是一些基于 binlog 监听的数据同步工具：



Canal 是一个由阿里巴巴开源的基于 MySQL 数据库 binlog 的增量订阅和消费组件，它主要用于 MySQL 数据变更捕获（Change Data Capture, CDC）。Canal 模拟 MySQL slave 的交互协议，将自己伪装成一个 MySQL slave 连接到 MySQL master 上，从而可以读取 binlog 并解析数据。

1. **Debezium**:
   Debezium 是一个非常成熟和完整的 CDC 解决方案，支持多种数据库系统，包括 MySQL、PostgreSQL、MongoDB 和 SQL Server。它提供了丰富的配置选项和扩展性，并且与 Kafka Connect 紧密集成。

2. **Maxwell's Daemon (Maxwell)**:
   Maxwell 也是一个成熟的工具，它专注于将 MySQL binlog 事件作为 JSON 输出到 Kafka。Maxwell 提供了一些配置选项，如数据过滤和转换，以及对多种消息队列的支持。

3. **Canal**:
   Canal 提供了与 MySQL binlog 的直接集成，并且可以将数据变更输出到多种消息队列，如 Kafka、RocketMQ 等。Canal 的特点是易于部署和使用，同时也支持集群模式和高可用性。

4. **go-mysql**:
   go-mysql 是一个更轻量级的库，它提供了基础的 binlog 解析功能。虽然它可能没有像 Debezium 或 Maxwell 那样的高级特性，但它允许开发者在 Go 语言环境中构建定制的数据同步或变更捕获解决方案。

5. **Tungsten Replicator**:
   Tungsten Replicator 是一个功能丰富的数据复制引擎，支持多种数据库系统。它提供了复杂的数据过滤和转换功能，以及跨多个数据库的复制能力。

6. **gh-ost**:
   gh-ost 主要用于 MySQL 的在线 schema 变更，但它也可以访问 binlog。虽然它不是一个传统的 CDC 工具，但它在进行 schema 变更时能够同步数据。

7. **Ripple**:
   Ripple 是一个较新的 MySQL binlog 服务器，它可以作为 MySQL 的中间层来复制 binlog 数据。Ripple 相对较新，可能不如其他工具成熟。

8. **Zongji**:
   Zongji 是一个 Node.js 库，用于监听 MySQL binlog 事件。它适合于需要在 Node.js 应用程序中实现实时数据同步的场景。

这些工具各有特点，适用于不同的场景和需求。在选择合适的工具时，你需要考虑你的特定用例、技术栈、性能要求以及是否需要与特定的消息队列或数据存储系统集成。



## Canal 集群

> 要启动多实例可以用canal的集群模式

Canal 集群是 Canal 的一种部署模式，旨在提供高可用性和负载均衡。在集群模式下，多个 Canal 实例（通常称为 Canal server）协同工作，由一个或多个协调节点（称为 Canal manager 或 Canal server manager）管理。这种架构可以提高系统的稳定性和扩展性，确保数据同步任务在面对单点故障或负载增加时仍能正常运行。

Canal 集群的关键组件包括：

1. **Canal Server**：
   这是集群中的工作节点，负责连接到 MySQL 服务器，监听 binlog 事件，并将变更数据推送到下游的消费者，如消息队列（Kafka、RocketMQ 等）。

2. **Canal Instance**：
   在每个 Canal server 中，可以运行一个或多个 Canal instance。每个 instance 对应于一个 MySQL 数据源，并独立进行数据同步。

3. **Canal Manager**：
   这是集群的协调节点，负责管理多个 Canal server 的状态和配置。Canal manager 会监控每个 server 的健康状况，并在需要时进行故障转移或重新分配任务。

4. **ZooKeeper**：
   Canal 集群通常使用 ZooKeeper 作为协调服务，用于管理集群状态、选举和分布式锁等。ZooKeeper 保证了集群中各个组件的一致性和协调。

Canal 集群的工作流程大致如下：

- 当集群启动时，Canal manager 会从 ZooKeeper 获取集群的配置信息，并根据这些信息启动和管理 Canal server。
- 每个 Canal server 连接到指定的 MySQL 数据源，监听 binlog 事件，并将数据变更推送到配置的下游消费者。
- 如果某个 Canal server 发生故障，Canal manager 会检测到这一情况，并通过 ZooKeeper 触发故障转移机制，将故障节点的任务重新分配给其他健康的 server。
- ZooKeeper 作为协调中心，确保集群中的 Canal server 之间不会发生冲突，并且在进行故障转移和任务重新分配时能够保持一致性。

使用 Canal 集群模式可以提高数据同步任务的可靠性和可伸缩性，但同时也增加了部署和运维的复杂性。因此，在决定是否使用集群模式时，需要根据实际的业务需求和技术能力进行权衡。