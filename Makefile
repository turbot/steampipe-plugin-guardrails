STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo
install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/guardrails@latest/steampipe-plugin-guardrails.plugin -tags "${BUILD_TAGS}" *.go
