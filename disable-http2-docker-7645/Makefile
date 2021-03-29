start:
	cd certs && mkcert whoami && cd -
	docker-compose up -d

clean:
	rm -rf certs/*

test:
	curl -kvv https://whoami

stop:
	docker-compose down --remove-orphans
