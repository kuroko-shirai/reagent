#/bin/bash
# Use Bash instead of SH
export SHELL := /bin/bash

.DEFAULT_GOAL := reagent

SERVICE       := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
FILENAME      := ${SERVICE}.proto

ABS_DIR_ROOT  := $(shell pwd)
ABS_PROTO_DIR := ${ABS_DIR_ROOT}/protofiles
ABS_FILEPATH  := ${ABS_PROTO_DIR}/${SERVICE}

REL_DIR_ROOT  := .
REL_PROTO_DIR := ${REL_DIR_ROOT}/protofiles
REL_FILEPATH  := ${REL_PROTO_DIR}/${SERVICE}

generate:
	rm -rf ${ABS_FILEPATH}/generated/
	docker build -t reagent-image \
	--build-arg ABS_FILEPATH="${ABS_FILEPATH}" \
	--build-arg REL_FILEPATH="${REL_FILEPATH}" \
	--build-arg FILENAME="${FILENAME}" \
	${ABS_DIR_ROOT}
	docker run -d --name reagent-container -it reagent-image sh
	docker cp reagent-container:/go/generated/ ${ABS_FILEPATH}
	docker stop reagent-container
	docker rm reagent-container
	docker rmi reagent-image

%::
	@true
