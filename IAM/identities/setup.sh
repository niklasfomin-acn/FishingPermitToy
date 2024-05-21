#!/bin/bash

# Run the keycloak server
docker run --name mykeycloak -p 8080:8080 \
        -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin \
        quay.io/keycloak/keycloak:latest \
        start-dev

# Get admin token with client credentials of admin-cli
export access_token=$(curl --insecure -X POST http://localhost:8080/realms/master/protocol/openid-connect/token --user admin-cli:lcM2yJh4J6jdUqFzUTDZonE7Rephl5mX -H 'content-type: application/x-www-form-urlencoded' -d 'username=admin&password=admin&grant_type=password' | jq --raw-output '.access_token' )

# Create another realm on Keycloak
curl -X POST -k -g -H "Authorization: Bearer $access_token" "http://localhost:8080/admin/realms" -H "Content-Type: application/json" --data '{"id": "testrealm","realm": "testrealm","accessTokenLifespan": 600,"enabled": true,"sslRequired": "all","bruteForceProtected": true,"loginTheme": "keycloak","eventsEnabled": false,"adminEventsEnabled": false}'

# Create users in the realm
curl -k -X POST http://localhost:8080/admin/realms/testrealm/users -H "Content-Type: application/json" -H "Authorization: Bearer $access_token" --data '{ "username": "Admin", "enabled": true, "realmRoles": [ "user", "offline_access" ], "attributes": { "uid": ["4010"], "homedir": ["/home/Admin"], "shell": ["/sbin/nologin"] } }'

# Get the Users IDs to set passwords
curl -k -X GET http://localhost:8080/admin/realms/testrealm/users -H "Authorization: Bearer "$access_token | jq

# Set password for the users
curl -k -X PUT http://localhost:8080/admin/realms/testrealm/users/<user-ID>/reset-password -H "Content-Type: application/json" -H "Authorization: bearer $access_token" --data '{ "type": "password", "temporary": false, "value": "test" }'

# Test Fetch a token for the user
# Remember to: Check Token Expiration Time, Disable SSL Verification and set up account fully for the user.
curl -k -X POST http://localhost:8080/realms/testrealm/protocol/openid-connect/token -d 'client_id=admin-cli' -d 'client_secret=wMHsoqeRmJGEGhgZulNmCrG8PiWTTXXc' -d 'username=niklas' -d 'password=test' -d 'grant_type=password' -H 'Content-Type: application/x-www-form-urlencoded'