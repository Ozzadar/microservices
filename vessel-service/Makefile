build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/ozzadar/microservices/vessel-service proto/vessel/vessel.proto
	sudo docker build -t vessel-service .

run:
	sudo docker run \
	-p 50052:50051 -e MICRO_SERVER_ADDRESS=:5051 \
	--network host vessel-service