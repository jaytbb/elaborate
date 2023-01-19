all: build run

build:
	cd server; go build

run:
	server/server

clean:
	rm -f server/server
