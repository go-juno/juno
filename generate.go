package main

//go:generate wire cmd/wire.go
//go:generate apidoc -i ./api  -o ./static/api-doc  -c ./docs --single
