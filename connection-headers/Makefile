start:
	docker-compose up

test:
	curl -vv -h "Connection: x-forwarded-for" http://whoami

stop:
	docker-compose down --remove-orphans
