#!/bin/bash

if docker-compose --version &> /dev/null; then
  echo -n "docker-compose";
elif docker compose version &> /dev/null; then
  echo -n "docker compose";
else
  exit 1;
fi
