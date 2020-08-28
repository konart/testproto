gen:
	protoc -I . -I `go list -m -f {{.Dir}} github.com/alta/protopatch` -I `go list -m -f {{.Dir}} google.golang.org/protobuf` --go-patch_out=plugin=go:. *.proto

run:
	go run main.go
