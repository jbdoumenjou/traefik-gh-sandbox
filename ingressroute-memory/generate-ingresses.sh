#!/usr/bin/bash

rm urls.txt
rm resources/*

for i in {000..100}
do
  cat ingressroute-template.yml | sed "s/%NUMBER%/-$i/" > ./resources/ingressroute-$i.yml
  printf "GET http://whoami.localhost/whoami-$i\n" >> urls.txt
done
