run:
	go run cmd/avyaas/main.go

ws:
	go run cmd/cli/ws.go

tidy:
	go mod tidy

migrate:
	go run cmd/cli/migrations.go

create_admin:
	go run cmd/cli/create_administrator.go

lint:
	golangci-lint run

transfer-data:
	go run cmd/cli/transfer_data.go

create-package-types:
	go run cmd/cli/create_package_types.go