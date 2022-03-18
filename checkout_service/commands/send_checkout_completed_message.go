package commands

import (
	"checkoutservice/infra"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type sendCheckoutCompletedMessage struct {
}

func NewSendCheckoutCompletedMessage() *sendCheckoutCompletedMessage {
	return &sendCheckoutCompletedMessage{}
}

func (s *sendCheckoutCompletedMessage) SendMessage(ctx context.Context, message string) error {
	span, _ := tracer.SpanFromContext(ctx)
	span = tracer.StartSpan("send_checkout_completed_message.send_message", tracer.ChildOf(span.Context()))

	sqsClient := infra.NewSQSClient()
	queueUrl := "http://localhost:4566/000000000000/checkout-completed-queue"
	span.SetTag("sqs_url", queueUrl)
	span.SetTag("message", message)

	carrier := tracer.TextMapCarrier{}
	tracer.Inject(span.Context(), carrier)

	carrierAsJSON, err := json.Marshal(carrier)

	if err != nil {
		span.Finish(tracer.WithError(err))
		return err
	}

	fmt.Println(string(carrierAsJSON))

	_, err = sqsClient.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"TraceCarrier": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(carrierAsJSON)),
			},
		},
		QueueUrl: &queueUrl,
	})

	if err != nil {
		span.Finish(tracer.WithError(err))
		return err
	}

	span.Finish()
	return nil
}
