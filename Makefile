
all: compose

reqs:
	@protoc --version || ( echo 'Please install protocolbuffers from  https://github.com/protocolbuffers/protobuf/releases'; exit 1)
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get github.com/micro/protoc-gen-micro

protobuf: reqs
	find . -maxdepth 3 -name '*.proto' -exec protoc --micro_out=. --go_out=. {} \;

build:
	find cmd -maxdepth 1 -type d -mindepth 1 -exec go build -o {}/$$(echo {} | sed 's:^.*/::') verbio/{} \; 

compose: protobuf
	find cmd -maxdepth 1 -mindepth 1 -type d -exec sh -c 'cd $$1 && CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-extldflags -static"' -- {} \;
	docker-compose up --build

test:
	go test verbio/nlu

		