.PHONY: build publish stop run package
IMAGES = bowling_api
BINARY_FOLDER_NAME = bin

# API
API_PACKAGE_FILE = api.tgz
API_FOLDER_LOCATION = ./api/src

all: build run

local:
	@cd ./api/src && go run .

# Gather all the go files and put them into a tar file so that it can be used in docker building
package:
	@cd $(API_FOLDER_LOCATION) && mkdir $(BINARY_FOLDER_NAME) && tar -czvf $(BINARY_FOLDER_NAME)/$(API_PACKAGE_FILE) `find . -name "*.go"`

build:
	@cd $(API_FOLDER_LOCATION) && docker image build -f ../../docker/Dockerfile.$(IMAGES) -t enf3rno/$(IMAGES):latest .

run:
	@cd ./docker && docker compose up -d

stop:
	@cd ./docker && docker compose down

publish:
	@docker login
	@$(foreach IMAGE,$(IMAGES), \
		docker image push enf3rno/$(IMAGE):latest ; \
	)

