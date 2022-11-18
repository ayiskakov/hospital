LOCAL_BIN:=$(CURDIR)/bin

.PHONY: dbup
dbup:
	docker compose up -d

.PHONY: dbdown
dbdown:
	docker compose down

.PHONY: start
start:
	go run cmd/srv/main.go