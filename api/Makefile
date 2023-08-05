ENTITY ?= $(error Please set ENTITY)
MIG_NAME ?= $(error Please set MIG_NAME)
MIG_VERSION ?= $(error Please set MIG_VERSION)

DB_USER ?= api_admin
DB_PASS ?= admin312
DB_URL ?= localhost:28032
DB_SOURCE = postgres://$(DB_USER):$(DB_PASS)@$(DB_URL)/bloated_api?search_path=public&sslmode=disable

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: run
run: ## Run server.
	go run ./cmd/...

.PHONY: mock
mock: mockgen ## Generate mock implementations.
	go generate ./...

.PHONY: test
test: ginkgo ## Run tests.
	$(GINKGO) ./...

##@ Schema

.PHONY: ent-new
ent-new: ## Create new entity.
	go run entgo.io/ent/cmd/ent new --target pkg/core/ent/schema $(ENTITY)

.PHONY: ent-gen
ent-gen: ## Generate Go code for entities.
	go run entgo.io/ent/cmd/ent generate --feature schema/snapshot ./pkg/core/ent/schema

.PHONY: ent-desc
ent-desc: ## Describe schemas.
	go run entgo.io/ent/cmd/ent describe ./pkg/core/ent/schema

.PHONY: mig-gen
mig-gen: atlas ## Generate migration from schema.
	atlas migrate diff $(MIG_NAME) \
	--dir "file://pkg/core/ent/migrate/migrations" \
	--to "ent://pkg/core/ent/schema" \
	--dev-url "docker://postgres/15/test?search_path=public"

.PHONY: mig-hash
mig-hash: atlas ## Re-hash checksum of migrations.
	atlas migrate hash \
	--dir "file://pkg/core/ent/migrate/migrations"

.PHONY: mig-status
mig-status: atlas ## Show status of migrationed version.
	atlas migrate status \
	--dir "file://pkg/core/ent/migrate/migrations" \
	--url "$(DB_SOURCE)"

.PHONY: mig-apply
mig-apply: atlas ## Apply migrations to DB.
	atlas migrate apply \
	--dir "file://pkg/core/ent/migrate/migrations" \
	--url "$(DB_SOURCE)"

.PHONY: mig-apply-to
mig-apply-to: atlas ## Migrate to specific version.
	atlas schema apply \
	--to "file://pkg/core/ent/migrate/migrations?version=$(MIG_VERSION)" \
	--dev-url "docker://postgres/15/test?search_path=public" \
	--exclude "atlas_schema_revisions" \
	--url "$(DB_SOURCE)"

	atlas migrate set $(MIG_VERSION) \
	--dir "file://pkg/core/ent/migrate/migrations" \
	--url "$(DB_SOURCE)"

.PHONY: mig-apply-dry
mig-apply-dry: atlas ## Dry run migrations to DB.
	atlas migrate apply \
	--dir "file://pkg/core/ent/migrate/migrations" \
	--url "$(DB_SOURCE)" \
	--dry-run

##@ Tool Dependencies

## Location to install dependencies to
LOCALBIN := $(shell pwd)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

## Tool Binaries
GINKGO ?= $(LOCALBIN)/ginkgo

## Tool Versions
GINKGO_VERSION ?= v2.11.0

.PHONY: ginkgo
ginkgo: $(GINKGO) ## Install ginkgo locally if necessary.
$(GINKGO): $(LOCALBIN)
	@test -s $(LOCALBIN)/ginkgo || GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_VERSION)

.PHONY: mockgen
mockgen: ## Install mockgen if necessary.
	@command -v mockgen > /dev/null || go install github.com/golang/mock/mockgen@latest

.PHONY: atlas
atlas: ## Install atlas if necessary.
	@command -v atlas > /dev/null || curl -sSf https://atlasgo.sh | sh
