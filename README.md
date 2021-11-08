# What is it?

We implement a simple SQL generator using golang.

It's nothing but a homework to learn golang interface.

# What does it look like?

run `make build` to build a binary file in `examples/`, and then run it:

```bash
# 03-sqlgenerator on 🌱 main via 🐹 v1.17.2 
❯ make build
go build -o examples/sql-generator.out examples/main.go

# 03-sqlgenerator on 🌱 main [!?] via 🐹 v1.17.2 
❯ ./examples/sql-generator.out 
SELECT [ID],[Name],[AGE] FROM [Student] WHERE [ID]='0906'  ORDER BY [ID] ASC 
SELECT [ID],[Name],[AGE] FROM [Student] WHERE [ID]='0906'  ORDER BY [ID] ASC  offset 0 rows fetch next 20 rows only 
SELECT `ID`,`Name`,`AGE` FROM `Student` WHERE `ID`='0906'  ORDER BY `ID` ASC 
SELECT `ID`,`Name`,`AGE` FROM `Student` WHERE `ID`='0906'  ORDER BY `ID` ASC  limit 0,20

```

# Referneces

This project is inspired by chapter 7 of [this book](http://product.dangdang.com/29120162.html).

---
