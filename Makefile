
DEBUG_PORT = 8888
CROSS_HOST = 172.2.0.21

build:
	rm -rf build
	mkdir -p build/echo-service/conf
	cp -f echo-service/cmd/echo-server/conf/config.yaml build/echo-service/conf/config.yaml
	go build -o build/echo-service/echo-server  echo-service/cmd/echo-server/main.go
	echo "build ok"

cpu_torch:
	rm -fr cpu.svg && go-torch -u http://$(CROSS_HOST):$(DEBUG_PORT) -t30 -p > cpu.svg

cpu_pprof:
	go tool pprof -http=:8081 --seconds 30 http://$(CROSS_HOST):$(DEBUG_PORT)

mem_torch:
	rm -fr mem.svg && go-torch -alloc_space http://$(CROSS_HOST):$(DEBUG_PORT)/debug/pprof/heap --colors=mem -p > mem.svg

mem_pprof:
	go tool pprof -inuse_space -http=:8081 http://$(CROSS_HOST):$(DEBUG_PORT)/debug/pprof/heap

trace:
	wget -O trace.out http://$(CROSS_HOST):$(DEBUG_PORT)/debug/pprof/trace && go tool trace trace.out

clean:
	rm -fr build cpu.svg mem.svg
