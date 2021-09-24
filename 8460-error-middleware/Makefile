start: stop
	docker-compose up -d

test:
	curl -X POST -d '500' http://whoami.localhost/health
	curl -svv whoami.localhost/health | grep Host:

stop:
	docker-compose down

clean: stop
	docker system prune -f
