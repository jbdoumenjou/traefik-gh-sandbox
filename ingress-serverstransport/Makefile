start:
	echo "Starting stack"
	k3d cluster delete mycluster
	k3d cluster create mycluster --api-port 6550 -p 80:80@loadbalancer  -p 8080:8080@loadbalancer -p 443:443@loadbalancer --k3s-server-arg '--no-deploy=traefik' -i rancher/k3s:v1.20.7-k3s1
 	# import the last compiled version of Traefik
	k3d image import traefik/traefik:latest -c mycluster
	# to ensure the definitions are loaded first
	# then load all other configurations
	kubectl apply -f conf/
	watch kubectl get all

reload:
	echo "Reload resources"
	kubectl apply -f conf/

stop:
	echo "Stopping stack"
	k3d cluster delete mycluster
	docker system prune -f
