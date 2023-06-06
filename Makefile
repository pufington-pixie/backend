SERVER_NAME := projectapi
PID_FILE := process.pid

.PHONY: start stop

start:
	@echo "Starting $(SERVER_NAME)..."
	@go run main.go & echo $$! > $(PID_FILE)

stop:
	@echo "Stopping $(SERVER_NAME)..."
	@-kill `cat $(PID_FILE)` 2> /dev/null || true
	@-rm $(PID_FILE) 2> /dev/null || true
