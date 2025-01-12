package broker

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type SaveTopic struct {
	Name string
}

type Storage interface {
	SaveTopic(topic SaveTopic) error
}

type DatabaseStorage struct {
	DB *sql.DB
}

func (storage *DatabaseStorage) SaveTopic(topic SaveTopic) error {
	fmt.Println("Hello from storage")

	return nil
}

type Topic struct {
	Name string
}

type Message struct {
	Topic   string `json:"topic"`
	Payload string `json:"payload"`
}

type Subscriber struct {
	Topic       Topic
	CallbackUrl string `json:"callback_url"`
}

type Broker struct {
	topics      map[string][]*Topic
	subscribers map[string][]*Subscriber
	messages    chan Message
	storage     Storage
}

func NewBroker(bufferSize int, storage Storage) *Broker {

	return &Broker{
		topics:      make(map[string][]*Topic),
		subscribers: make(map[string][]*Subscriber),
		messages:    make(chan Message, bufferSize),
		storage:     storage,
	}
}

func (b *Broker) Publish(msg Message) {
	b.messages <- msg
	fmt.Printf("Message published to topic %s\n", msg.Topic)
}

func (b *Broker) Subscribe(sub *Subscriber) {
	_, topicExists := b.topics[sub.Topic.Name]
	if !topicExists {
		fmt.Printf("Oppsi, topic %s does not exists", sub.Topic.Name)
		return
	}

	subscribers := b.subscribers[sub.Topic.Name]
	for _, subb := range subscribers {
		if subb.CallbackUrl == sub.CallbackUrl {
			fmt.Printf("Already subscriber to the topic %s", sub.Topic.Name)
			return
		}

	}

	b.subscribers[sub.Topic.Name] = append(b.subscribers[sub.Topic.Name], sub)
}

func (b *Broker) CreateTopic(topic string) {

	b.storage.SaveTopic(SaveTopic{Name: "Hello there"})

	_, exists := b.topics[topic]
	if !exists {
		b.topics[topic] = append(b.topics[topic], &Topic{Name: topic})
	}
}

// func (b *Broker) GetTopics() map[]*Topic {
// 	topics := make([]Topic, 0)
// 	for _, topicList := range b.topics {
// 		for _, topic := range topicList {
// 			topics = append(topics, *topic)
// 		}
// 	}

// 	return topics
// }

func (b *Broker) StartWorkers(workerCount int, client *http.Client) {
	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			fmt.Printf("Worker %d started\n", workerId)

			for msg := range b.messages {
				b.dispatchMessage(workerId, msg, client)
			}
		}(i)
	}
}

func (b *Broker) dispatchMessage(workerId int, msg Message, client *http.Client) {
	// b.mu.RLock()
	// defer b.mu.RUnlock()

	fmt.Printf("Mesasge for topic name: %s, payload: %s, worker %d\n", msg.Topic, msg.Payload, workerId)

	subscribers, exists := b.subscribers[msg.Topic]
	if !exists {
		fmt.Printf("No subscribers for topic %s\n", msg.Topic)
		return
	}

	for _, sub := range subscribers {
		go process(workerId, msg, *sub, client)

	}

}

func process(workerId int, msg Message, sub Subscriber, client *http.Client) {

	maxRetries := 3
	initialDelay := 500 * time.Millisecond
	maxDelay := 3 * time.Second

	requestPayload := map[string]string{"payload": msg.Payload}
	toJson, err := json.Marshal(requestPayload)

	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	requestBody := bytes.NewBuffer(toJson)

	callback := func() error {
		req, req_err := http.NewRequest("POST", sub.CallbackUrl+fmt.Sprintf("?workerId=%d", workerId), requestBody)
		if req_err != nil {
			return req_err
		}

		res, res_err := client.Do(req)
		if res_err != nil {
			return res_err
		}

		defer res.Body.Close()

		fmt.Printf("Status code: %d\n", res.StatusCode)

		if res.StatusCode >= http.StatusBadRequest {
			return fmt.Errorf("Http error: %d", res.StatusCode)
		}

		return nil
	}

	backoff_err := backoff(callback, maxRetries, initialDelay, maxDelay)
	if backoff_err != nil {
		fmt.Println("Final error:", backoff_err)
	} else {
		fmt.Println("Operation succeeded!")
	}

}

func backoff(callback func() error, maxRetries int, initialDelay time.Duration, maxDelay time.Duration) error {
	delay := initialDelay

	for i := 0; i < maxRetries; i++ {
		err := callback()

		if err == nil {
			return nil
		}

		fmt.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, delay)

		time.Sleep(delay)

		delay = time.Duration(float64(delay) * (1.5 + rand.Float64()*0.5)) // Add jitter
		if delay > maxDelay {
			delay = maxDelay
		}
	}

	return fmt.Errorf("Operation failed after %d retries", maxRetries)
}
