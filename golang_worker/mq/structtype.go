package mq

type RabbitMQQueue struct {
	Name       string `json:"name"`
	Messages   int    `json:"messages"`
	Durable    bool   `json:"durable"`
	Exclusive  bool   `json:"exclusive"`
	AutoDelete bool   `json:"auto_delete"`
}

type RabbitMQQueueDetail struct {
	Name                   string `json:"name"`
	MessagesReady          int    `json:"messages_ready"`
	MessagesUnacknowledged int    `json:"messages_unacknowledged"`
}

type MessageBody struct {
	Message string `json:"message"`
}
