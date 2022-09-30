APP_BUILD = build/gatewayApp

.PHONY: build start

build:
	clear; go build -o ${APP_BUILD} -v cmd/gatewayApp/main.go

start: build
	${APP_BUILD}