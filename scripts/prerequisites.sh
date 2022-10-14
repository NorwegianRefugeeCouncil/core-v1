#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)
red='\033[0;31m'
nc='\033[0m'

NEEDED_COMMANDS="envoy go"

for cmd in ${NEEDED_COMMANDS} ; do
    if ! command -v ${cmd} &> /dev/null ; then
        echo ${cmd} is not installed! Check README.md
        exit 1
    fi
done

if ! command -v envoy &> /dev/null ; then
echo -e ${red}Error!${nc}
echo -e Could not find ${bold}envoy${normal} on your PATH!
echo -e Install envoy ${bold}using brew install envoy${normal}
echo
fi

if ! docker-compose version &> /dev/null && ! docker compose version &> /dev/null; then
echo -e ${red}Error!${nc}
echo -e Could not find ${bold}docker-compose${normal} or ${bold}docker compose${normal} on your PATH!
echo -e https://docs.docker.com/get-docker/
echo
fi