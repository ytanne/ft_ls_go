OUTPUT_BIN = ft_ls

build:
	@echo "###\tBuilding binary executable"
	@go build -o $(OUTPUT_BIN) main.go

clear:
	@echo "###\tClearing the binary executable"
	@rm $(OUTPUT_BIN)