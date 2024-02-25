# Order Processing System
1. Order Receiving Module:
. Must receive order data in real time.
. Should send a confirmation back to the sender once the order is received.
2. Messaging Queue Integration (Apache Kafka):
. Implement Kafka producers to send order data to a Kafka topic.
. Ensure the system can handle high throughput (millions of orders).
3. Order Processing Module:
. Develop Kafka consumers to process orders from the Kafka topic.
. Orders should be processed and stored in the database sequentially.
. Ensure reliable processing (handling of failures, retries).
## Features
### Installation
git clone https://github.com/dungly116/order-processing.git

1.Install dependencies:
go mod tidy

2.Run the project:
go run cmd/main.go

