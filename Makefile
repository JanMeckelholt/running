PROTO_SRC_FILES := $(wildcard common/grpc/*/*.proto)
PROTO_OBJ_FILES := $(patsubst %.proto, %.pb.go, $(PROTO_SRC_FILES))
PROTO_OBJ_FILES_DART := $(patsubst %.proto, %.pbserver.dart, $(PROTO_SRC_FILES))
PROTO_OBJ_FILES_DART := $(patsubst common/%, running_app/lib/gen/%, $(PROTO_OBJ_FILES_DART))

.PHONY: generate_protos clean_protos generate_certs	get_clients
generate_protos: clean_protos $(PROTO_OBJ_FILES) generate_protos_dart
%.pb.go: %.proto
	protoc --proto_path common/grpc --go_opt=paths=source_relative --go_out=common/grpc --go-grpc_opt=paths=source_relative --go-grpc_out=common/grpc $<

generate_protos_dart: clean_protos_dart $(PROTO_OBJ_FILES_DART)
running_app/lib/gen/%.pbserver.dart: common/%.proto
	protoc --proto_path common/grpc --dart_out=grpc:running_app/lib/gen/grpc google/protobuf/empty.proto
	protoc --proto_path common/grpc --dart_out=grpc:running_app/lib/gen/grpc $<


clean_protos:
	rm -f $(PROTO_OBJ_FILES) 
	
clean_protos_dart:
	rm -f $(PROTO_OBJ_FILES_DART)


generate_certs:
	cd certs; ./gen_certs.sh; cd ..;

get_clients:
	docker exec -i postgres psql -U root running < ./get_clients.sql > db_clients.txt
	cat db_clients.txt