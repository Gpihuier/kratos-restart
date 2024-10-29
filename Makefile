PROJECT_NAME='kratos-restart'

.PHONY: api
# generate api
api:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd'

.PHONY: wire
# generate wire
wire:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'


.PHONY: docker-compose
docker-compose:
	cd deploy/docker-compose && docker-compose -f docker-compose.yml -p $(PROJECT_NAME) up -d