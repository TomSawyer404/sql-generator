debug: cmd/sql-generator/main.go
	go build -o cmd/sql-generator/debug.out -gcflags "-N -l" $^

build: cmd/sql-generator/main.go
	go build -o cmd/sql-generator/sql-generator.out $^

clean:
	rm cmd/sql-generator/*.out
