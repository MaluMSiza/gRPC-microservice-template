# Makefile para gRPC Microservice Template

.PHONY: generate build-server build-client clean test docker-build docker-run help

# Variáveis
PROTO_DIR := proto
GEN_DIR := gen
SERVER_BIN := grpc-server
CLIENT_BIN := grpc-client

# Gera os arquivos gRPC
generate:
	@mkdir -p $(GEN_DIR)
	protoc --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) $(PROTO_DIR)/greet.proto
	@echo "Arquivos gRPC gerados em $(GEN_DIR)/"

# Compila o servidor
build-server:
	@go build -o $(SERVER_BIN) ./cmd/server/main.go
	@echo "Servidor compilado: ./$(SERVER_BIN)"

# Compila o cliente
build-client:
	@go build -o $(CLIENT_BIN) ./cmd/client/main.go
	@echo "Cliente compilado: ./$(CLIENT_BIN)"

# Remove binários e arquivos gerados
clean:
	@rm -rf $(GEN_DIR) $(SERVER_BIN) $(CLIENT_BIN)
	@echo "Arquivos limpos!"

# Executa testes (adicione seus testes aqui)
test:
	@go test -v ./...

# Constrói a imagem Docker
docker-build:
	@docker-compose build

# Executa os serviços no Docker
docker-run:
	@docker-compose up

# Ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  generate       - Gera arquivos gRPC"
	@echo "  build-server   - Compila o servidor"
	@echo "  build-client   - Compila o cliente"
	@echo "  clean          - Remove arquivos gerados"
	@echo "  test           - Executa testes"
	@echo "  docker-build   - Constrói as imagens Docker"
	@echo "  docker-run     - Inicia os serviços no Docker"