build: accload

install:
	- go get github.com/tsenart/vegeta/lib
	- go get github.com/BurntSushi/toml

accload: *.go
	go build -o $@ $^
