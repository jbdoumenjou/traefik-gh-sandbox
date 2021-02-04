#! /bin/sh

case "$1" in
  certs)
    echo "Generate whoami certs"
    cd certs && rm -f * && mkcert "whoami" && mkcert "traefik.localhost" && cd -
  ;;
  test)
    echo "Test: try to reach whoami with tls v1.1"
    curl -kv --tls-max 1.1  https://whoami
  ;;
  clean)
    rm -f certs/*.pem
  ;;
esac