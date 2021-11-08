debug: examples/main.go
	go build -o examples/debug.out -gcflags "-N -l" $^

build: examples/main.go
	go build -o examples/sql-generator.out $^

clean:
	rm examples/*.out
