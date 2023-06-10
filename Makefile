.PHONY: build publish stop run package frontend bindir api apackage abuild
# VARS
FRONTEND_PACKAGE_FILE = frontend
API_PACKAGE_FILE = api
BINARY_FOLDER_NAME = bin
PACKAGES = $(API_PACKAGE_FILE) $(FRONTEND_PACKAGE_FILE)

all: build run
package: fpackage apackage
build: fbuild abuild

# Convienance functions
bindir:
	@$(foreach PACKAGE,$(PACKAGES), \
		cd ./$(PACKAGE) && mkdir -p $(BINARY_FOLDER_NAME) && cd - ; \
	)
clean:
	@rm -rf `find . -type d -name bin`
#END CONFIENANCE

# Frontend
frontend: fbuild run
fpackage: clean bindir
	@cd ./$(FRONTEND_PACKAGE_FILE)/src && tar -czvf ../$(BINARY_FOLDER_NAME)/$(FRONTEND_PACKAGE_FILE).tgz `find . -name "*.go"`
fbuild: fpackage
	@cd $(FRONTEND_PACKAGE_FILE) && docker image build -f ../docker/Dockerfile.bowling_$(FRONTEND_PACKAGE_FILE) -t enf3rno/bowling_$(FRONTEND_PACKAGE_FILE):latest .
## END FRONTEND

# API
api: abuild run
apackage: clean bindir
	@cd ./$(API_PACKAGE_FILE)/src && tar -czvf ../$(BINARY_FOLDER_NAME)/$(API_PACKAGE_FILE).tgz `find . -name "*.go"`
abuild: apackage
	@cd $(API_PACKAGE_FILE) && docker image build -f ../docker/Dockerfile.bowling_$(API_PACKAGE_FILE) -t enf3rno/bowling_$(API_PACKAGE_FILE):latest .
## END API

publish:
	@docker login
	@$(foreach IMAGE,$(IMAGES), \
		docker image push enf3rno/bowling_$(IMAGE):latest ; \
	)

# Compose functions
run:
	@cd ./docker && docker compose up

rund:
	@cd ./docker && docker compose up -d

stop:
	@cd ./docker && docker compose down
#

