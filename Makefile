include .env
DSN=host=${HOST} port=${DATABASE_PORT} user=${DATABASE_USERNAME} password=${DATABASE_PASSWORD} dbname=${DATABASE_NAME} sslmode=disable timezone=UTC connect_timeout=5
BINARY_NAME=subly.exe
REDIS="${REDIS_HOST}:${REDIS_PORT}"

## build: builds all binaries
build:
	@go build -o ${BINARY_NAME} ./cmd/web
	@echo back end built!

run: build
	@echo Starting...
	set "DSN=${DSN}"
	set "REDIS=${REDIS}"
	start /min cmd /c ${BINARY_NAME} &
	@echo back end started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

start: run

stop:
	@echo "Stopping..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped back end

restart: stop start

test:
	@echo "Testing..."
	go test -v ./...