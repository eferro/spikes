go run kafkatest.go -topic=test -verbose=true -brokers localhost:9092 -offset=oldest

go run kafkaproducer.go -key id1 -topic test -value data1


# Start local kafka
bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties


