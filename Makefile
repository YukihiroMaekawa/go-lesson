.PHONY: build

build:
	go build ./cmd/go-lesson

# gRPC-ドキュメント作成
gendoc:
	protoc --doc_out=./doc/gRPC --doc_opt=markdown,index.md ./proto/*.proto
