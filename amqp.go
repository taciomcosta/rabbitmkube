package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type queueResponse struct {
	Name          string `json:"name"`
	TotalMessages int    `json:"messages"`
}

// Queues represents a map of queue x messages on queue
type Queues map[string]int

// FetchQueues returns a pair of queues x messages on queue
func FetchQueues() Queues {
	return parseQueues(requestAMQPServer())
}

func requestAMQPServer() []queueResponse {
	client := &http.Client{}
	response, err := client.Do(newRequest())
	if err != nil {
		fmt.Println("Error while request amqp server")
		return []queueResponse{}
	}
	return parseResponse(response)
}

func newRequest() *http.Request {
	req, err := http.NewRequest("GET", os.Getenv("RABBITMQ_URL"), nil)
	if err != nil {
		fmt.Println("Error fecthing server")
	}
	req.SetBasicAuth(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"))
	return req
}

func parseResponse(response *http.Response) []queueResponse {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []queueResponse{}
	}
	return unmarshalResponse(bytes)
}

func unmarshalResponse(response []byte) []queueResponse {
	var queueResponses []queueResponse
	json.Unmarshal(response, &queueResponses)
	return queueResponses
}

func parseQueues(queueResponses []queueResponse) Queues {
	queues := make(Queues, 0)
	for _, qr := range queueResponses {
		queues[qr.Name] = qr.TotalMessages
	}
	return queues
}
