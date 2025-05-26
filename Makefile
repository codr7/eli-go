all:
	go build -C src -o ../bin/eli
	bin/eli

test: 
	go test src/tests/*

clean:
	rm -f bin/*
