# PHONY: generate
# generate:
# 	mkdir -p pkg/note_v1
# 	protoc	--proto_path api/note_v1 \
# 			--go_out=pkg/note_v1 --go_opt=paths=source_relative \
# 			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
# 			api/note_v1/note.proto

PHONY: generate
generate:
		mkdir -p pkg/note_v1
		protoc --proto_path vendor.protogen --proto_path api/note_v1 \
				--go_out=pkg/note_v1 --go_opt=paths=source_relative \
				--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=pkg/note_v1 \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=source_relative \
				--validate_out lang=go:pkg/note_v1\
				--swagger_out=allow_merge=true,merge_file_name=api:pkg/note_v1 \
				api/note_v1/note.proto
		mv pkg/note_v1/github.com/MuhahaSam/golangPractice/pkg/note_v1/* pkg/note_v1/
		rm -rf pkg/note_v1/github.com

server-start: go run .\cmd\server\main.go
client-start: go run .\cmd\client\main.go

LOCAL_MIGRATION_DIR = './migration'
LOCAL_MIGRATION_DNS = 'host=localhost port=5433 dbname=note user=postgres password=postgres'

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DNS} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DNS} down -v

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi