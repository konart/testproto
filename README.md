# install protobuf

brew install protobuf

# install patcher

go install github.com/alta/protopatch/cmd/protoc-gen-go-patch

# build proto and run

make gen &&  make run
