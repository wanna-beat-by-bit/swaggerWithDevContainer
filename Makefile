
.PHONY: run
run:
	go run cmd/swaggerTest/

.PHONY: dev-build
dev-build:
	docker compose --file docker-compose.dev.yml build

.PHONY: dev-start
dev-start:
	docker compose --file docker-compose.dev.yml up

.PHONY: help
help:
	# swag init - to generate swagger docs from project code

.PHONY: swag
swag:
	swag init -g cmd/swaggerTest/main.go	

.PHONY: vscode-post-install
vscode-post-install:
	# Пост установка плагинов для vs code для Fomin.AK
	code --install-extension usernamehw.errorlens
	code --install-extension golang.Go
	code --install-extension vscodevim.vim
	code --install-extension adpyke.codesnap
