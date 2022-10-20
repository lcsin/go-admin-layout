.PHONY: build
build:
	docker build -t goadmin-srv:latest -f Dockerfile .

.PHONY: dev test migrate
dev:
	go build
	go-admin.exe server -c config/settings-dev.yml
migrate:
	go build
	go-admin.exe migrate -c config/settings-dev.yml
