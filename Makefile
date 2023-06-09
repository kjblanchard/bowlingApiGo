.PHONY: build

all: build run

build:
	@cd ./src && docker image build -t enf3rno/bowling_api:latest .

run:
	@cd ./docker && docker compose up -d

stop:
	@cd ./docker && docker compose down
