PROTO_SRC_FILES := $(wildcard common/grpc/*/*.proto)
PROTO_OBJ_FILES := $(patsubst %.proto, %.pb.go, $(PROTO_SRC_FILES))

.PHONY: generate_protos clean_protos format-gorm
generate_protos: clean_protos $(PROTO_OBJ_FILES)

%.pb.go: %.proto
	protoc --proto_path common/grpc --go_opt=paths=source_relative --go_out=common/grpc --go-grpc_opt=paths=source_relative --go-grpc_out=common/grpc $<


clean_protos:
	rm -f $(PROTO_OBJ_FILES)

