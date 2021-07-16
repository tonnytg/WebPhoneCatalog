.ONESHELL:
.SHELL := /usr/bin/bash
CURRENT_FOLDER=$(shell basename "$$(pwd)")
ENV="dev"

local:
	docker-compose up --build -d