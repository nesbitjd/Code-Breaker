# The `clean` target is intended to clean the workspace
# and prepare the local changes for submission.
#
# Usage: `make clean`
.PHONY: clean
clean: tidy vet fmt 

# The `restart` target is intended to destroy and
# create the local Docker compose stack.
#
# Usage: `make restart`
.PHONY: restart
restart: down up

# The `up` target is intended to create
# the local Docker compose stack.
#
# Usage: `make up`
.PHONY: up
up: go-build compose-up

# The `down` target is intended to destroy
# the local Docker compose stack.
#
# Usage: `make down`
.PHONY: down
down: compose-down

# The `tidy` target is intended to clean up
# the Go module files (go.mod & go.sum).
#
# Usage: `make tidy`
.PHONY: tidy
tidy:
	@echo
	@echo "### Tidying Go module"
	@go mod tidy

# The `vet` target is intended to inspect the
# Go source code for potential issues.
#
# Usage: `make vet`
.PHONY: vet
vet:
	@echo
	@echo "### Vetting Go code"
	@go vet ./...

# The `fmt` target is intended to format the
# Go source code to meet the language standards.
#
# Usage: `make fmt`
.PHONY: fmt
fmt:
	@echo
	@echo "### Formatting Go Code"
	@go fmt ./...

.PHONY: go-build
go-build:
	@echo
	GOOS=linux CGO_ENABLED=0 go build cmd/coder breaker/*

.PHONY: docker-build
docker-build:
	@echo
	docker build -t registry.digitalocean.com/nesbitjd/codebreaker:latest .

# The `fix` target is intended to rewrite the
# Go source code using old APIs.
#
# Usage: `make fix`
.PHONY: fix
fix:
	@echo
	@echo "### Fixing Go Code"
	@go fix ./...

# The `pull` target is intended to pull all
# images for the local Docker compose stack.
#
# Usage: `make pull`
.PHONY: pull
pull:
	@echo
	@echo "### Pulling images for docker-compose stack"
	@docker-compose pull

#
# Usage: `make build-and-compose`
.PHONY: build-and-compose
build-and-compose:
	@echo
	@echo "### Creating and building containers for docker-compose stack"
	@docker-compose -f docker-compose.yml up -d --build

# The `compose-up` target is intended to build and create
# all containers for the local Docker compose stack.
#
# Usage: `make compose-up`
.PHONY: compose-up
compose-up:
	@echo
	@echo "### Creating containers for docker-compose stack"
	@docker-compose -f docker-compose.yml up -d 

# The `compose-down` target is intended to destroy
# all containers for the local Docker compose stack.
#
# Usage: `make compose-down`
.PHONY: compose-down
compose-down:
	@echo
	@echo "### Destroying containers for docker-compose stack"
	@docker-compose -f docker-compose.yml down