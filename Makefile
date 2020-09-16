.PHONY: help clean mk

SERVERNAME = echo

help :
	@echo "help info"

dist :
	rm -rf dist
	mkdir -p dist/$SERVERNAME-service/conf
	cp -f $SERVERNAME-service/cmd/$SERVERNAME-server/conf/config.yaml dist/$SERVERNAME-service/conf/config.yaml
	go build -o dist/$SERVERNAME-service/$SERVERNAME-server $SERVERNAME-service/cmd/$SERVERNAME-server/main.go

clean :
	@rm -rf dist