build: accload

install:
	- go get github.com/tsenart/vegeta/lib

accload: *.go
	go build -o $@ $^
