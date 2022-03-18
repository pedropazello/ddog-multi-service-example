
# sqs
createQueue(){
    local queueName=$1
    local deadLetterQueueName="${queueName}_deadletter"

    # create deadletterqueue
    awslocal sqs create-queue \
        --attributes '{
            "VisibilityTimeout": "30"
        }' \
        --queue-name $deadLetterQueueName

    # create queue
    awslocal sqs create-queue \
        --attributes '{
            "VisibilityTimeout": "30",
            "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:000000000000:'$deadLetterQueueName'\",\"maxReceiveCount\":\"5\"}"
        }' \
        --queue-name $queueName
}

listQueuees(){
    awslocal sqs list-queues
}

createQueue "checkout-completed-queue"
