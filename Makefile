.PHONY: local-test
local-test: clean
	docker-compose up --force-recreate

.PHONY: clean
clean:
	docker-compose down
