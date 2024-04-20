#!/bin/bash

# Run docker container in dev mode
docker run --name mykeycloak -p 8080:8080 \
        -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=<password> \
        quay.io/keycloak/keycloak:latest \
        start-dev