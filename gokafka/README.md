docker-compose -f docker-compose-single-broker.yml up -d

go run kafkatest.go -topic=test -verbose=true -brokers localhost:9092 -offset=oldest

go run kafkaproducer.go -key id1 -topic test -value data1
