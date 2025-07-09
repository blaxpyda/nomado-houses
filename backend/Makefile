include .env

.PHONY: help create-migration migrate-up migrate-down migrate-force

help: ## Show help
	@echo "\n\033[1mAvailable commands:\033[0m\n"
	@awk 'BEGIN {FS = ":.*##";} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

create-migration: ## Create a new migration file
	@read -p "Enter the sequence name: " SEQ; \
	    migrate create -ext sql -dir internal/database/migrations -seq $$SEQ

migrate-up: ## Run all up migrations
	@migrate -path internal/database/migrations -database "$(DB_URL)" up

migrate-down: ## Run all down migrations
	@migrate -path internal/database/migrations -database "$(DB_URL)" down

migrate-force: ## Force a migration to a specific version
	@read -p "Enter the version to force: " VERSION; \
	migrate -path internal/database/migrations -database "$(DB_URL)" force $$VERSION

