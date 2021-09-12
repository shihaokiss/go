package config

type AppConfig struct {
	KafakApp Kafak `ini:"kafak"`
	TailLogApp TailLog `ini:"taillog"`
}

type Kafak struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type TailLog struct {
	FailPath string `ini:"filepath"`
}