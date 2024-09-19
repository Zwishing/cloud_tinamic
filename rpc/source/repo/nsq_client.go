package repo

import (
	"cloud_tinamic/kitex_gen/data/source"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nsqio/go-nsq"
)

// MinioEvent defines the structure for MinIO events
type MinioEvent struct {
	EventName string `json:"EventName"`
	Key       string `json:"Key"`
	Size      int64  `json:"Size"`
	Bucket    string `json:"Bucket"`
}

// NSQHandler defines the NSQ message handler
type NSQHandler struct {
	SourceRepo SourceRepo
}

// NewNSQHandler creates a new NSQ handler
func NewNSQHandler(repo SourceRepo) nsq.Handler {
	return &NSQHandler{SourceRepo: repo}
}

// HandleMessage processes received NSQ messages
func (h *NSQHandler) HandleMessage(m *nsq.Message) error {
	var event MinioEvent
	if err := json.Unmarshal(m.Body, &event); err != nil {
		klog.Errorf("Failed to unmarshal NSQ message: %v", err)
		return err
	}

	klog.Infof("Received MinIO event: %s", event.Key)
	item := &source.Item{
		Name:         event.Key,
		ItemType:     0, // Determine the appropriate ItemType
		Key:          event.Key,
		Size:         event.Size,
		ModifiedTime: "", // Set the appropriate ModifiedTime
		Path:         event.Bucket + "/" + event.Key,
	}

	// Insert file metadata into the database
	if _, err := h.SourceRepo.AddItem(1, event.Bucket, item); err != nil {
		klog.Errorf("Failed to insert file metadata into database: %v", err)
		return err
	}
	return nil
}

// NsqConsumer wraps an NSQ consumer
type NsqConsumer struct {
	Consumer *nsq.Consumer
}

// NewNsqConsumer creates a new NSQ consumer
func NewNsqConsumer(config *NsqConfig, handler nsq.Handler) (*NsqConsumer, error) {
	consumer, err := nsq.NewConsumer(config.Topic, config.Channel, nsq.NewConfig())
	if err != nil {
		klog.Errorf("Failed to create NSQ consumer for topic %s, channel %s: %v", config.Topic, config.Channel, err)
		return nil, err
	}

	consumer.AddHandler(handler)

	if err := consumer.ConnectToNSQLookupd(config.Address); err != nil {
		consumer.Stop()
		klog.Errorf("Failed to connect to NSQLookupd at %s for topic %s, channel %s: %v", config.Address, config.Topic, config.Channel, err)
		return nil, err
	}

	klog.Infof("Successfully created NSQ consumer for topic %s, channel %s, and connected to NSQLookupd at %s", config.Topic, config.Channel, config.Address)

	return &NsqConsumer{Consumer: consumer}, nil
}

// Stop gracefully stops the NSQ consumer
func (c *NsqConsumer) Stop() {
	c.Consumer.Stop()
}
