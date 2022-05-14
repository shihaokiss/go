package conf

type Kafka struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type TailLog struct {
	FileName string `ini:"filename"`
}

type AppConf struct {
	Kafka `ini:"kafka"`
	TailLog `ini:"taillog"`
}
