PHONY: generate
generate:
	mkdir -p pkg/note_v1
	protoc	--proto_path api/note_v1 \
			--go_out=pkg/note_v1 --go_opt=paths=source_relative \
			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
			api/note_v1/note.proto

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