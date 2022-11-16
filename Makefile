LOCAL_BIN:=$(CURDIR)/bin

.PHONY: dbup
dbup:
	docker compose up -d

.PHONY: dbdown
dbdown:
	docker compose down

.PHONE: server
server:
	go run cmd/srv/main.go