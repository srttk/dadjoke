targets := $(wildcard *.go) $(wildcard **/*.go)
dadjoke:$(targets)
	go build
install: dadjoke
	cp ./dadjoke /usr/local/bin
	chmod +x /usr/local/bin
