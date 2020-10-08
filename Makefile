export GOFLAGS := -mod=vendor
export GO111MODULE := on

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags timetzdata
	sudo docker build -t variadico/hello-nats-dapr:1 .
	sudo docker push variadico/hello-nats-dapr:1
