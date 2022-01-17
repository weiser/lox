# lox
implementation of Lox in Crafting Compilers by Nystrom

## How to build interpreter

```
go build -o lox
```

## How to start the interpreter:

```
./lox
```

## How to add to the grammar of the language:

1. modify `cmd/tool/generateAst.go` to reflect the changes to the grammar
2. run `go run cmd/tool/generateAst.go expr`