APP_NAME := gtfs-validator
MAIN_PKG := ./
ENV_FILE ?= ../enviromments/production/stops/.env
STOPS_REFERENCE_API_URL ?= http://localhost:3001/stops
STOPS_REFERENCE_API_TIMEOUT ?= 5s
API_HEALTH_URL ?= http://localhost:3001/health
API_START_CMD ?= cd ../ts-wrapper && npm run dev:api

.PHONY: all build run dev test clean help

all: build

build:
	go build -o ./../bin/$(APP_NAME) $(MAIN_PKG)

run: build
	./../bin/$(APP_NAME) $(ARGS)

dev:
	@set -e; \
	ENV_FILE_PATH="$(ENV_FILE)"; \
	if [ ! -f "$$ENV_FILE_PATH" ]; then \
		ENV_FILE_PATH="../enviromments/production/stops/env.example"; \
	fi; \
	if [ -f "$$ENV_FILE_PATH" ]; then \
		echo "Loading env vars from $$ENV_FILE_PATH"; \
		set -a; . "$$ENV_FILE_PATH"; set +a; \
	else \
		echo "No env file found (.env or env.example), using Makefile defaults"; \
	fi; \
	API_START_CMD_VAL="$${API_START_CMD:-$(API_START_CMD)}"; \
	STOPS_REFERENCE_API_URL_VAL="$${STOPS_REFERENCE_API_URL:-$(STOPS_REFERENCE_API_URL)}"; \
	STOPS_REFERENCE_API_TIMEOUT_VAL="$${STOPS_REFERENCE_API_TIMEOUT:-$(STOPS_REFERENCE_API_TIMEOUT)}"; \
	API_PORT_VAL="$${API_PORT:-$$(printf '%s\n' "$$STOPS_REFERENCE_API_URL_VAL" | sed -n 's#.*://[^:/]*:\([0-9][0-9]*\)/.*#\1#p')}"; \
	if [ -z "$$API_PORT_VAL" ]; then API_PORT_VAL=3000; fi; \
	API_STARTED=0; \
	API_PID=""; \
	CURL_CMD="curl -fsS"; \
	API_URL_BASE="$$(printf '%s\n' "$$STOPS_REFERENCE_API_URL_VAL" | sed -n 's#\(.*://[^:/]*:[0-9][0-9]*\).*#\1#p')"; \
	if [ -z "$$API_URL_BASE" ]; then API_URL_BASE="http://localhost:$$API_PORT_VAL"; fi; \
	API_HEALTH_CHECK_URL="$${API_HEALTH_URL:-$$API_URL_BASE/health}"; \
	cleanup() { \
		if [ "$$API_STARTED" = "1" ] && [ -n "$$API_PID" ]; then \
			echo "Stopping API (PID=$$API_PID)..."; \
			kill $$API_PID >/dev/null 2>&1 || true; \
		fi; \
	}; \
	if $$CURL_CMD "$$STOPS_REFERENCE_API_URL_VAL" >/dev/null 2>&1; then \
		echo "API already running and reference endpoint is reachable at $$STOPS_REFERENCE_API_URL_VAL"; \
	else \
		echo "Starting API on port $$API_PORT_VAL..."; \
		FOUND_FREE_PORT=0; \
		CURRENT_API_PORT=$$API_PORT_VAL; \
		for p_offset in $$(seq 0 20); do \
			TEST_PORT=$$((CURRENT_API_PORT + p_offset)); \
			if ! lsof -i :$$TEST_PORT >/dev/null 2>&1; then \
				API_PORT_VAL=$$TEST_PORT; \
				echo "Found free port: $$API_PORT_VAL"; \
				FOUND_FREE_PORT=1; \
				break; \
			fi; \
		done; \
		if [ $$FOUND_FREE_PORT -eq 0 ]; then \
			echo "Configured API port $$CURRENT_API_PORT is already in use and no free port found in range +20."; \
			exit 1; \
		fi; \
		STOPS_REFERENCE_API_URL_VAL="$$(printf '%s\n' "$$STOPS_REFERENCE_API_URL_VAL" | sed "s#:[0-9][0-9]*/#:$$API_PORT_VAL/#")"; \
		API_HEALTH_CHECK_URL="$$(printf '%s\n' "$$API_HEALTH_CHECK_URL" | sed "s#:[0-9][0-9]*/#:$$API_PORT_VAL/#")"; \
		PORT="$$API_PORT_VAL" sh -c "$$API_START_CMD_VAL" >/tmp/gtfs-api.log 2>&1 & \
		API_PID=$$!; \
		API_STARTED=1; \
		trap cleanup EXIT INT TERM; \
		echo "Waiting reference endpoint: $$STOPS_REFERENCE_API_URL_VAL"; \
		for i in $$(seq 1 30); do \
			if $$CURL_CMD "$$STOPS_REFERENCE_API_URL_VAL" >/dev/null 2>&1; then \
				echo "API is ready"; \
				break; \
			fi; \
			if $$CURL_CMD "$$API_HEALTH_CHECK_URL" >/dev/null 2>&1; then \
				echo "API is ready (health check ok; reference endpoint not ready)"; \
				break; \
			fi; \
			if ! kill -0 $$API_PID >/dev/null 2>&1; then \
				echo "API process exited before exposing the reference endpoint."; \
				echo "---- /tmp/gtfs-api.log ----"; \
				tail -n 80 /tmp/gtfs-api.log || true; \
				echo "---------------------------"; \
				cleanup; \
				exit 1; \
			fi; \
			sleep 1; \
			if [ $$i -eq 30 ]; then \
				echo "API did not expose reference endpoint in time"; \
				cleanup; \
				exit 1; \
			fi; \
		done; \
	fi; \
	trap cleanup EXIT INT TERM; \
	DATABASE_URI="$${DATABASE_URI:-}" \
	STOPS_REFERENCE_API_URL="$$STOPS_REFERENCE_API_URL_VAL" \
	STOPS_REFERENCE_API_TIMEOUT="$$STOPS_REFERENCE_API_TIMEOUT_VAL" \
	go run main.go -out=output.json -rules=../tml-rules.json -lang=pt -log=debug --input="$(INPUT)"
test:
	go test ./...

clean:
	rm -rf bin

help:
	@echo "Available targets:"
	@echo "  make build        Build the validator binary to bin/$(APP_NAME)"
	@echo "  make run ARGS='...'  Build and run with extra CLI arguments"
	@echo "  make dev INPUT='...' [ENV_FILE='../enviromments/production/stops/.env']"
	@echo "  make test         Run Go tests"
	@echo "  make clean        Remove build artifacts"