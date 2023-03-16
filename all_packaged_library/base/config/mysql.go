package config

type IMySQLConf interface {
	GetDriverName() string
	GetDataSourceName() string
	GetMaxOpenConn() int
	GetMaxIdleConn() int
	GetMaxConnLifeTime() int
}

// MySQLConf mysql配置
type MySQLConf struct {
	DriverName      string `yaml:"driver_name" mapstructure:"driver_name"`
	DataSourceName  string `yaml:"data_source_name" mapstructure:"data_source_name"`
	MaxOpenConn     int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleConn     int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `yaml:"max_conn_life_time" mapstructure:"max_conn_life_time"`
}

func (m MySQLConf) GetDriverName() string {
	return m.DriverName
}

func (m MySQLConf) GetDataSourceName() string {
	return m.DataSourceName
}

func (m MySQLConf) GetMaxOpenConn() int {
	return m.MaxOpenConn
}

func (m MySQLConf) GetMaxIdleConn() int {
	return m.MaxIdleConn
}

func (m MySQLConf) GetMaxConnLifeTime() int {
	return m.MaxConnLifeTime
}
