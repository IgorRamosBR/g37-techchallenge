package clients

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sqsQueue struct {
	sqs      *sqs.Client
	queueURL string
}

func NewSQSQueue(sqs *sqs.Client, queueURL string) Queue {
	return sqsQueue{
		sqs:      sqs,
		queueURL: queueURL,
	}
}
func (q sqsQueue) SendMessage(data []byte) error {
	_, err := q.sqs.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody: aws.String(string(data)),
		QueueUrl:    aws.String(q.queueURL),
	})
	if err != nil {
		return fmt.Errorf("failed to send message to the queue, error: %v", err)
	}

	return nil
}
