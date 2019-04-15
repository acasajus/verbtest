
all: reqs protobuf

reqs:
	@protoc --version || ( echo 'Please install protocolbuffers from  https://github.com/protocolbuffers/protobuf/releases'; exit 1)
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get github.com/micro/protoc-gen-micro

protobuf: reqs
	find . -name '*.proto' -maxdepth 3 -exec protoc --micro_out=. --go_out=. {} \;

build:
	find cmd -maxdepth 1 -type d -mindepth 1 -exec go build -o {}/$$(echo {} | sed 's:^.*/::') verbio/{} \; 

compose: protobuf
	find cmd -maxdepth 1 -type d -mindepth 1 -exec sh -c 'cd $$1 && GOOS=linux go build' -- {} \;
	docker-compose up --build

		