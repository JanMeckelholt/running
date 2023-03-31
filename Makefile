PROTO_SRC_FILES := $(wildcard common/grpc/*/*.proto)
PROTO_OBJ_FILES := $(patsubst %.proto, %.pb.go, $(PROTO_SRC_FILES))

.PHONY: generate_protos clean_protos generate_certs	get_clients
generate_protos: clean_protos $(PROTO_OBJ_FILES)

%.pb.go: %.proto
	protoc --proto_path common/grpc --go_opt=paths=source_relative --go_out=common/grpc --go-grpc_opt=paths=source_relative --go-grpc_out=common/grpc $<


clean_protos:
	rm -f $(PROTO_OBJ_FILES)


generate_certs:
	cd certs; ./gen_certs.sh; cd ..;

get_clients:
	docker exec -i postgres psql -U root running < ./get_clients.sql > db_clients.txt
	cat db_clients.txt