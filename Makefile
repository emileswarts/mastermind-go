build: stop
	docker-compose build

stop: 
	docker-compose down

test: build
	docker-compose exec -T app "go test -v ./..."

serve: build
	docker-compose run app sh -c "sleep infinity"

run: 
	docker-compose up

shell: 
	docker-compose run app bash

.PHONY: build test
