up: down
	docker-compose up

down:
	docker-compose down --remove-orphans

test:
	# should show X-Forwarded-Tls-Client-Cert-Info: Subject%3D%22OU%3DSimple+Signing+Section%2COU%3DSimple+Signing+Section+2%22 in response
	curl -kv https://whoami.localhost --cert ./certs/cheese.org.crt --key ./certs/cheese.org.key
