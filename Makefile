run:
	go run ./cmd/gateway

generate:
	@echo "🧩 Generating inbound clients from contracts..."
	go run ./pkg/bus/cmd/main.go
	@echo "✅ Done."