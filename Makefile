.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  build                  Compile the Terraform provider binary"
	@echo "  install-plugin         Install compiled provider binary to local Terraform plugin directory"
	@echo "  init                   Initialize Terraform in each examples/* project"
	@echo "  validate               Validate Terraform configuration in each project"
	@echo "  fmt-check              Check formatting of Terraform files"
	@echo "  fmt                    Format Terraform files"
	@echo "  docs                   Generate terraform-docs in each project (if main.tf exists)"
	@echo "  o-init                 Initialize OpenTofu in each examples/* project"
	@echo "  o-validate             Validate OpenTofu configuration"
	@echo "  o-fmt-check            Check formatting of OpenTofu files"
	@echo "  o-fmt                  Format OpenTofu files"
	@echo "  up                     Start Docker Compose services"
	@echo "  launch                 Open https://localhost:5100 in default browser"
	@echo "  down                   Stop Docker Compose services"
	@echo "  go-fmt-check           Check formatting of Go source files"
	@echo "  go-fmt                 Format Go source files"
	@echo ""
	@echo "Environment:"
	@echo "  TDIR                   Directory to run Terraform/OpenTofu in (set internally)"
	@echo "  TCMD                   Terraform/OpenTofu command (init, validate, fmt, etc.)"
	@echo ""

### Terraform
.PHONY: build
build:
	go build -o terraform-provider-pwpush

.PHONY: install-plugin
install-plugin:
	mkdir -p ~/.terraform.d/plugins/localdomain/local/pwpush/0.1.0/linux_amd64/
	cp terraform-provider-pwpush ~/.terraform.d/plugins/localdomain/local/pwpush/0.1.0/linux_amd64/

.PHONY: init
init:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _terraform TDIR=$$project TCMD=init ; \
	done

.PHONY: validate
validate:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _terraform TDIR=$$project TCMD=validate ; \
	done

.PHONY: fmt-check
fmt-check:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _terraform TDIR=$$project TCMD="fmt -check" ; \
	done

.PHONY: fmt
fmt:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _terraform TDIR=$$project TCMD="fmt" ; \
	done
	@for project in $$(find e2e-tests -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _terraform TDIR=$$project TCMD="fmt" ; \
	done

_terraform:
	terraform -chdir=${TDIR} ${TCMD}

### DOCS
.PHONY: docs
docs:
	@for dir in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		if [ -f $$dir/main.tf ]; then \
			terraform-docs -c .terraform-docs.yml $$dir; \
		fi; \
	done
	@for dir in $$(find e2e-tests -type d -mindepth 1 -maxdepth 1); do \
		if [ -f $$dir/main.tf ]; then \
			terraform-docs -c .terraform-docs.yml $$dir; \
		fi; \
	done

### Opentofu
.PHONY: o-init
o-init:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _opentofu TDIR=$$project TCMD=init ; \
	done

.PHONY: o-validate
o-validate:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _opentofu TDIR=$$project TCMD=validate ; \
	done

.PHONY: o-fmt-check
o-fmt-check:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _opentofu TDIR=$$project TCMD="fmt -check" ; \
	done

.PHONY: o-fmt
o-fmt:
	@for project in $$(find examples -type d -mindepth 1 -maxdepth 1); do \
		$(MAKE) _opentofu TDIR=$$project TCMD="fmt" ; \
	done

_opentofu:
	tofu -chdir=${TDIR} ${TCMD}

### Docker
.PHONY: up
up:
	docker compose -f docker/docker-compose.yml up -d

.PHONY: launch
launch:
	@PWPUSHER_HOST=$${PWPUSHER_HOST:-'localhost:5100'} ; \
	URL=$${URL:-http://$${PWPUSHER_HOST}} ; \
	echo "Opening $${URL} ..." ; \
	OS=$$(uname | tr '[:upper:]' '[:lower:]') ; \
	if [ "$$OS" = "linux" ]; then \
		xdg-open "$${URL}" >/dev/null 2>&1 || echo "Could not open browser (xdg-open not found?)" ; \
	elif [ "$$OS" = "darwin" ]; then \
		open "$${URL}" ; \
	elif echo "$$OS" | grep -q "mingw\\|msys\\|cygwin"; then \
		start "$${URL}" ; \
	else \
		echo "Cannot open browser automatically on this OS: $$OS" ; \
	fi

.PHONY: down
down:
	docker compose -f docker/docker-compose.yml down

### Go
.PHONY: go-fmt-check
go-fmt-check:
	@echo "Checking Go code formatting..."
	@unformatted=$$(find . -type f -name '*.go' -not -path './pwpush_data/*' 2>/dev/null | xargs gofmt -s -l); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not properly formatted:"; \
		echo "$$unformatted"; \
		echo ""; \
		echo "Run 'make go-fmt' to format them."; \
		exit 1; \
	else \
		echo "All Go files are properly formatted."; \
	fi

.PHONY: go-fmt
go-fmt:
	@echo "Formatting Go code..."
	@find . -type f -name '*.go' -not -path './pwpush_data/*' 2>/dev/null | xargs gofmt -s -w
	@echo "Done."