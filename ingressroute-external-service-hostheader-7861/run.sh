#! /bin/sh
case "$1" in
  start)
    echo "Starting stack"
#    cd certs/ && mkcert whoami && cd -
    k3d cluster delete mycluster
    k3d cluster create mycluster --api-port 6550 -p 80:80@loadbalancer  -p 8080:8080@loadbalancer -p 443:443@loadbalancer --k3s-server-arg '--no-deploy=traefik' -i rancher/k3s:v1.18.6-k3s1
    # to ensure the definitions are loaded first
    # then load all other configurations
    kubectl apply -f conf/
  ;;
  reload)
    echo "Reload resources"
    kubectl apply -f conf/
  ;;
  stop)
    echo "Stopping stack"
    k3d cluster delete mycluster
  ;;

esac
