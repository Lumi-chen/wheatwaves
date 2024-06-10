package config

// 创建一个config的结构体
// 属性于config.yaml里对应
type Config struct {
	MySql  MySql  `yaml:"mysql"`
	System System `yaml:"system"`
	Logger Logger `yaml:"logger"`
}
