package kafak

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
    client sarama.SyncProducer
	err error
)

func Init(connStr []string) (err error) {
	// 设置配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll                      // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner             // 新选出一个分区
	config.Producer.Return.Successes = true                               // 成功交付的消息在 success chan 返回

    // 连接kafka
	client, err = sarama.NewSyncProducer(connStr, config)
	if err != nil {
		fmt.Println("connect kafka err: ", err)
		return
	}
	fmt.Println("Connect success.")
	return
}

func SendMessage(topic string, data string) (err error) {
	// 构造发送消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic                                                 // 设置消息主题
	msg.Value = sarama.StringEncoder(data)

	// 发送日志
	paritition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg err: ", err)
		return
	}
	fmt.Println(paritition, offset)
	return
}

