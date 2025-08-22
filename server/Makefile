include .env
export $(shell sed 's/=.*//' .env)

.PHONY: mup
mup:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database '$(DB_ADDR)' -path internal/migrations up


.PHONY: mdown
mdown:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database '$(DB_ADDR)' -path internal/migrations down

.PHONY: mupprod
mup:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database '$(DB_ADDR)' -path internal/migrations up

.PHONY: mdownprod
mdown:
	@echo "Running DB migration on: $(DB_ADDR)"
	migrate -database '$(DB_ADDR)' -path internal/migrations down

.PHONY: dkup
dkup:
	@echo "Starting Docker containers..."
	docker compose up

.PHONY: dkdown
dkdown:
	@echo "Stopping Docker containers..."
	docker compose down