package repo

import (
	"cloud_tinamic/kitex_gen/data/source"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nsqio/go-nsq"
)

// MinioEvent 定义 MinIO 事件结构
type MinioEvent struct {
	EventName string `json:"EventName"`
	Key       string `json:"Key"`
	Size      int64  `json:"Size"`
	Bucket    string `json:"Bucket"`
}

// NSQHandler 定义 NSQ 消息处理器
type NSQHandler struct {
	SourceRepo SourceRepo
}

func NewNSQHandler(repo SourceRepo) nsq.Handler {
	return &NSQHandler{SourceRepo: repo}
}

// HandleMessage 处理接收到的 NSQ 消息
func (h *NSQHandler) HandleMessage(m *nsq.Message) error {
	var event MinioEvent
	err := json.Unmarshal(m.Body, &event)
	if err != nil {
		klog.Errorf("Failed to unmarshal NSQ message: %v", err)
		return err
	}

	klog.Infof("Received MinIO event: %s", event.Key)
	item := &source.Item{
		Name:         "",
		ItemType:     0,
		Key:          "",
		Size:         0,
		ModifiedTime: "",
		Path:         "",
	}

	// 插入文件元数据到数据库
	_, err = h.SourceRepo.AddItem(1, "", item)
	if err != nil {
		return err
	}
	if err != nil {
		klog.Errorf("Failed to insert file metadata into database: %v", err)
		return err
	}
	return nil
}

type NsqConsumer struct {
	Consumer *nsq.Consumer
}

func NewNsqConsumer(config *NsqConfig, handler nsq.Handler) *NsqConsumer {
	consumer, err := nsq.NewConsumer(config.Topic, config.Channel, nsq.NewConfig())
	if err != nil {
		klog.Errorf("Failed to create NSQ consumer for topic %s, channel %s: %v", config.Topic, config.Channel, err)
		return nil
	}

	consumer.AddHandler(handler)

	if err := consumer.ConnectToNSQLookupd(config.Address); err != nil {
		klog.Errorf("Failed to connect to NSQLookupd at %s for topic %s, channel %s: %v", config.Address, config.Topic, config.Channel, err)
		return nil
	}

	klog.Infof("Successfully created NSQ consumer for topic %s, channel %s, and connected to NSQLookupd at %s", config.Topic, config.Channel, config.Address)

	return &NsqConsumer{Consumer: consumer}
}

func (c *NsqConsumer) Stop() {
	c.Consumer.Stop()
}
