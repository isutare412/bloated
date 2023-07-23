TAG_UI ?= latest
ENV_FILE ?= .env

IMAGE_UI = bloated/ui:$(TAG_UI)

COMPOSE_CMD = docker-compose -f docker-compose.yaml --env-file $(ENV_FILE) -p bloated

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: ui
ui: ## Build docker image of UI.
	docker build -f ui/Dockerfile -t $(IMAGE_UI) ui

##@ Deployment

.PHONY: up
up: ## Run components.
	$(COMPOSE_CMD) up -d

.PHONY: down
down: ## Shutdown components.
	$(COMPOSE_CMD) down

.PHONY: p
ps: ## Print running components.
	$(COMPOSE_CMD) ps

.PHONY: log
logs: ## Tail logs of components.
	$(COMPOSE_CMD) logs -f