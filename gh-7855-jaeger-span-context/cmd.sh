#! /bin/sh

case "$1" in
  certs)
    echo "Generate whoami certs"
    cd certs && rm -f * && mkcert "whoami" && mkcert "traefik.localhost" && cd -
  ;;
  clean)
    rm -f certs/*.pem
  ;;
esac