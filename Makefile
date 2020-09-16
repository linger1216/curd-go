.PHONY: cross help clean mk dist cpu_torch cpu_pprof mem_torch mem_pprof trace doc

SERVER_NAME = echo
DEBUG_PORT = 21211
HOST = localhost
SCP_HOST = 114.67.106.133
SCP_PATH = /root/projects

help :
	@echo "dist -> 编译"
	@echo "cpu_torch -> CPU 火焰图"
	@echo "cpu_pprof -> CPU pprof"
	@echo "mem_torch -> MEM 火焰图"
	@echo "mem_pprof -> MEM pprof"
	@echo "trace -> 追踪"
	@echo "cross -> scp部署"


run :
	cd $(SERVER_NAME)-service/cmd/$(SERVER_NAME)-server && go run main.go

dist :
	mkdir -p dist/$(SERVER_NAME)-service/conf
	cp -f $(SERVER_NAME)-service/cmd/$(SERVER_NAME)-server/conf/config.yaml dist/$(SERVER_NAME)-service/conf/config.yaml
	go build -o dist/$(SERVER_NAME)-service/$(SERVER_NAME)-server $(SERVER_NAME)-service/cmd/$(SERVER_NAME)-server/main.go

cpu_torch :
	mkdir -p dist/$(SERVER_NAME)-service/conf
	go-torch -u http://$(HOST):$(DEBUG_PORT) -t30 -p > dist/cpu.svg

cpu_pprof :
	go tool pprof -http=:8081 --seconds 30 http://$(HOST):$(DEBUG_PORT)

mem_torch :
	mkdir -p dist/$(SERVER_NAME)-service/conf
	go-torch -alloc_space http://$(HOST):$(DEBUG_PORT)/debug/pprof/heap --colors=mem -p > dist/mem.svg

mem_pprof :
	go tool pprof -inuse_space -http=:8081 http://$(HOST):$(DEBUG_PORT)/debug/pprof/heap

trace :
	mkdir -p dist/$(SERVER_NAME)-service/conf
	wget -O trace.out http://$(HOST):$(DEBUG_PORT)/debug/pprof/trace && go tool trace dist/trace.out

cross :
	mkdir -p dist/$(SERVER_NAME)-service/conf
	cp -f $(SERVER_NAME)-service/cmd/$(SERVER_NAME)-server/conf/config_ec2.yaml dist/$(SERVER_NAME)-service/conf/config.yaml
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/$(SERVER_NAME)-service/$(SERVER_NAME)-server $(SERVER_NAME)-service/cmd/$(SERVER_NAME)-server/main.go
	scp -r dist/* root@$(SCP_HOST):$(SCP_PATH)

doc :
	apidoc -i $(SERVER_NAME)-service/handlers

clean :
	@rm -rf dist
	@rm -rf doc