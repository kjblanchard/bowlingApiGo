.PHONY: build publish stop run package frontend bindir
#  Frontend
FRONTEND_PACKAGE_FILE = frontend
# API
API_PACKAGE_FILE = api
# VARS
BINARY_FOLDER_NAME = bin
PACKAGES = $(API_PACKAGE_FILE) $(FRONTEND_PACKAGE_FILE)

all: clean package build run

package: fpackage apackage

build: fbuild abuild

bindir:
	@$(foreach PACKAGE,$(PACKAGES), \
		cd ./$(PACKAGE) && mkdir -p $(BINARY_FOLDER_NAME) && cd - ; \
	)

clean:
	@rm -rf `find . -type d -name bin`

local:
	@cd ./api/src && go run .
# Frontend
frontend: clean fpackage fbuild

fpackage: bindir
	@cd ./$(FRONTEND_PACKAGE_FILE)/src && tar -czvf ../$(BINARY_FOLDER_NAME)/$(FRONTEND_PACKAGE_FILE).tgz `find . -name "*.go"`

flocal:
	@cd ./frontend/src && go run .

fbuild:
	@cd $(FRONTEND_PACKAGE_FILE) && docker image build -f ../docker/Dockerfile.bowling_$(FRONTEND_PACKAGE_FILE) -t enf3rno/bowling_$(FRONTEND_PACKAGE_FILE):latest .

apackage:
	@cd $(API_FOLDER_LOCATION) && mkdir -p $(BINARY_FOLDER_NAME) && tar -czvf $(BINARY_FOLDER_NAME)/$(API_PACKAGE_FILE).tgz `find . -name "*.go"`


abuild:
	@cd $(API_FOLDER_LOCATION) && docker image build -f ../../docker/Dockerfile.bowling_$(API_PACKAGE_FILE) -t enf3rno/bowling_$(API_PACKAGE_FILE):latest .

publish:
	@docker login
	@$(foreach IMAGE,$(IMAGES), \
		docker image push enf3rno/bowling_$(IMAGE):latest ; \
	)

# Compose functions
run:
	@cd ./docker && docker compose up

rund:
	@cd ./docker && docker compose up

stop:
	@cd ./docker && docker compose down
#

