SHELL:= /bin/bash

.DEFAULT_GOAL := init

.PHONY: init
init: init.dir 

.PHONY: build
build: api.server comment.server

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(shell pwd)
endif

ifeq ($(origin DATA_DIR),undefined)
DATA_DIR := $(ROOT_DIR)/data
endif

ifeq ($(origin API_SRV_DIR),undefined)
API_SRV_DIR := $(ROOT_DIR)/cmd/api
endif

ifeq ($(origin COMMENT_SRV_DIR),undefined)
COMMENT_SRV_DIR := $(ROOT_DIR)/cmd/comment
endif

ifeq ($(origin FAVORITE_SRV_DIR),undefined)
FAVORITE_SRV_DIR := $(ROOT_DIR)/cmd/favorite
endif

ifeq ($(origin FEED_SRV_DIR),undefined)
FEED_SRV_DIR := $(ROOT_DIR)/cmd/feed
endif

ifeq ($(origin PUBLISH_SRV_DIR),undefined)
PUBLISH_SRV_DIR := $(ROOT_DIR)/cmd/publish
endif

ifeq ($(origin RELATION_SRV_DIR),undefined)
RELATION_SRV_DIR := $(ROOT_DIR)/cmd/relation
endif

ifeq ($(origin USER_SRV_DIR),undefined)
USER_SRV_DIR := $(ROOT_DIR)/cmd/user
endif


.PHONY: init.dir
init.dir:
	@echo "===========> Run init.dir"
	@echo "ROOT_DIR: "$(ROOT_DIR) 
	@echo "DATA_DIR: "$(DATA_DIR) 
	

.PHONY: minio.server
minio.server:
	@echo "===========> Run minio.server" 
	@echo $(DATA_DIR)
	@minio server $(DATA_DIR)

.PHONY: etcd.server
etcd.server:
	@echo "===========> Run etcd.server"
	@etcd

.PHONY: api.server
api.server:
	@echo "===========> Run api.server" 
	@cd $(API_SRV_DIR);go run api.go

.PHONY: comment.server
comment.server:
	@echo "===========> Run comment.server" 
	@cd $(COMMENT_SRV_DIR);go run comment.go
	

.PHONY: favorite.server
favorite.server:
	@echo "===========> Run favorite.server" 
	@cd $(FAVORITE_SRV_DIR);go run favorite.go

.PHONY: feed.server
feed.server:
	@echo "===========> Run feed.server" 
	@cd $(FEED_SRV_DIR);go run feed.go

.PHONY: publish.server
publish.server:
	@echo "===========> Run publish.server" 
	@cd $(PUBLISH_SRV_DIR);go run publish.go

.PHONY: relation.server
relation.server:
	@echo "===========> Run relation.server" 
	@cd $(RELATION_SRV_DIR);go run relation.go

.PHONY: user.server
user.server:
	@echo "===========> Run user.server" 
	@cd $(USER_SRV_DIR);go run user.go
