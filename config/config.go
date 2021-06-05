package config

const (
	// ==== 数据源配置信息 =====
	DriverName = "mysql"
	Username   = "root"
	Password   = "toor"
	Host       = "localhost"
	Port       = "3306"
	DataBase   = "ginhello"
	// 最大连接数
	MaxOpenConns = 10
	// 最大空闲连接数
	MaxIdleConns = 1

	// ==== JWT 配置信息 =====
	// Secret 密钥
	Secret = "gin hello"
	// OneDayOfHours 设定过期时间为一天
	OneDayOfHours = 60 * 60 * 24
)