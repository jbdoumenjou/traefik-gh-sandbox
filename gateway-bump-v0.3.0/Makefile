all: start

mkcert:
	go run ${GOROOT}/src/crypto/tls/generate_cert.go  --rsa-bits 4096 --host whoami.localhost --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
	mv cert.pem certs/whoami.pem
	mv key.pem certs/whoami-key.pem

start:
	echo "Starting stack"
	k3d cluster delete mycluster
	k3d cluster create mycluster \
		--api-port 6550 \
		-p 80:80@loadbalancer \
		-p 443:443@loadbalancer \
		--k3s-server-arg '--no-deploy=traefik'
	k3d image import traefik/traefik:latest -c mycluster
	kubectl create secret tls mysecret --cert certs/whoami.pem --key certs/whoami-key.pem
	kubectl apply -f gateway-resources
	kubectl apply -f services/whoami-svc.yml

apply-all: apply-http apply-tcp apply-tls

apply-http:
	kubectl apply -f simple-http

apply-tcp:
	kubectl apply -f simple-tcp

apply-tls:
	kubectl apply -f simple-tls

test:
	(echo WHO; cat) | openssl s_client -connect whoami.localhost:9443

reload:
	echo "Reload resources"
	kubectl apply -f gateway-resources
	kubectl apply -f services

stop:
	echo "Stopping stack"
	k3d cluster delete mycluster

gstatus:
	kubectl get gateways -o json | jq ".items[] | .metadata.name , .status"

clean:
	$(RM) -f certs/whoami.pem certs/whoami-key.pem
