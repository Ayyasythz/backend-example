dependency:
	@echo ">> Downloading Dependencies"
	@go mod download

swag-init:
	@echo ">> Running swagger init"
	@swag init

run-api: dependency swag-init
	@echo ">> Running API Server"
	@go run main.go serve-http

test:
	@echo ">> Running Test"
	@go test -v -cover -count=1 -failfast ./...

remock:
	#https://github.com/vektra/mockery
	@echo ">> Mock Repositories"
	@mockery --all --dir ./internal/domain/repository --output ./internal/domain/repository/mocks_repository --outpkg mocks_repository

	@echo ">> Mock UseCases"
	@mockery --all --dir ./internal/usecases --output ./internal/usecases/mocks_usecases --outpkg mocks_usecases

	@echo ">> Mock API Middleware"
	@mockery --all --dir ./internal/handler/api/middleware --output ./internal/handler/api/middleware/mocks_middleware --outpkg mocks_middleware

	@echo ">> Mock Interfaces"
	@mockery --all --recursive --dir ./internal/interfaces --output ./internal/interfaces/mocks_interfaces --outpkg mocks_interfaces

	@echo ">> Mock Infra"
	@mockery --all --recursive --dir ./internal/infrastructures --output ./internal/infrastructures/mocks_infrastructures --outpkg mocks_infrastructures