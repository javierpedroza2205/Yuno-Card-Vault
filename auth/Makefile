# Variables
PROTO_DIR = proto
OUT_DIR = $(PROTO_DIR)
PROTO_FILE = $(PROTO_DIR)/auth.proto
GOPATH_BIN = $(shell go env GOPATH)/bin

# Comandos de instalación
PROTOC_GEN_GO = $(GOPATH_BIN)/protoc-gen-go
PROTOC_GEN_MICRO = $(GOPATH_BIN)/protoc-gen-micro

all: proto

# Generar los archivos .go a partir de auth.proto
proto: $(PROTO_FILE) $(PROTOC_GEN_GO) $(PROTOC_GEN_MICRO)
	@mkdir -p $(OUT_DIR)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--micro_out=$(OUT_DIR) --micro_opt=paths=source_relative \
		$(PROTO_FILE)

# Instala el plugin de Go para Protobuf si no está instalado
$(PROTOC_GEN_GO):
	@echo "Installing protoc-gen-go..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


# Instala el plugin de Go para Micro si no está instalado
$(PROTOC_GEN_MICRO):
	@echo "Installing protoc-gen-micro..."
	go install github.com/micro/protoc-gen-micro@latest

# Limpia los archivos generados
clean:
	@echo "Cleaning generated files..."
	@rm -rf $(OUT_DIR)

# Ejecuta limpieza seguida de compilación
rebuild: clean proto

.PHONY: all proto clean rebuild
