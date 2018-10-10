build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/ozzadar/microservices/user-service proto/user/user.proto
	sudo docker build -t user-service .

run:
	sudo docker run \
	-p 50053:50051 -e MICRO_SERVER_ADDRESS=:5051 \
	--network host vessel-service