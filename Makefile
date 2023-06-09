.PHONY: build publish stop run package frontend
IMAGES = api frontend
BINARY_FOLDER_NAME = bin

#  Frontend
FRONTEND_PACKAGE_FILE = frontend
FRONTEND_FOLDER_LOCATION = ./frontend/src

# API
API_PACKAGE_FILE = api
API_FOLDER_LOCATION = ./api/src

all: clean package build run

local:
	@cd ./api/src && go run .

frontend:
	@cd ./frontend/src && go run .

package: fpackage apackage

build: fbuild abuild

apackage:
	@cd $(API_FOLDER_LOCATION) && mkdir -p $(BINARY_FOLDER_NAME) && tar -czvf $(BINARY_FOLDER_NAME)/$(API_PACKAGE_FILE).tgz `find . -name "*.go"`

fpackage:
	@cd $(FRONTEND_FOLDER_LOCATION) && mkdir -p $(BINARY_FOLDER_NAME) && tar -czvf $(BINARY_FOLDER_NAME)/$(FRONTEND_PACKAGE_FILE).tgz `find . -name "*.go"`

abuild:
	@cd $(API_FOLDER_LOCATION) && docker image build -f ../../docker/Dockerfile.bowling_$(API_PACKAGE_FILE) -t enf3rno/bowling_$(API_PACKAGE_FILE):latest .

fbuild:
	@cd $(FRONTEND_FOLDER_LOCATION) && docker image build -f ../../docker/Dockerfile.bowling_$(FRONTEND_PACKAGE_FILE) -t enf3rno/bowling_$(FRONTEND_PACKAGE_FILE):latest .

publish:
	@docker login
	@$(foreach IMAGE,$(IMAGES), \
		docker image push enf3rno/bowling_$(IMAGE):latest ; \
	)
clean:
	@rm -rf $(find . -type d -name bin)

# Compose functions
run:
	@cd ./docker && docker compose up

rund:
	@cd ./docker && docker compose up

stop:
	@cd ./docker && docker compose down
#

