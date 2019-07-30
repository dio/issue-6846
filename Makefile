TAG ?= grpctranscodingexample:v1

up:
	make -C docker/dev up

upd:
	make -C docker/dev upd

down:
	make -C docker/dev down

docker-build:
	docker build -t $(TAG) .

.PHONY: up upd down docker-build
