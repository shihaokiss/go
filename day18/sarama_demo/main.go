package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	// 设置配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll                      // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner             // 新选出一个分区
	config.Producer.Return.Successes = true                               // 成功交付的消息在 success chan 返回

    // 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("connect kafka err: ", err)
		return
	}
	defer client.Close()                                                  // 函数退出关闭连接

	// 构造发送消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"                                                 // 设置消息主题
	msg.Value = sarama.StringEncoder("this is web log.")

	// 发送日志
	paritition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg err: ", err)
	}
	fmt.Println(paritition, offset)
}