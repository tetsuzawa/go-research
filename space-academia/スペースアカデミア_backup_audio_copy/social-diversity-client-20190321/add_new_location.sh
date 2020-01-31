#!/bin/bash

#server_address="http://localhost:8080"
server_address="https://social-diversity-devel.appspot.com"

observer_file="observer.json"
location_file="location.json"

wait_time=1

email=$(cat ${observer_file} | jq -r '.["e-mail"]')
password=$(cat ${observer_file} | jq -r '.password')

auth_success_token=$(curl -X POST -H "Content-Type: application/json" -d "{\"e-mail\": \"${email}\", \"password\": \"${password}\"}" ${server_address}/login | jq -r ".auth_success_token" )

location_id=$(curl -X POST -H "Authorization: Bearer ${auth_success_token}" -H "Content-Type: application/json" -d @${location_file}  ${server_address}/locations | jq -r ".location_id")

if [ "${location_id}" = "" ]; then
    echo "Adding the new location failed."
    exit 1
fi

saved_filename="${location_file/./_${location_id}.}"

cp ${location_file} ${saved_filename}

echo "Your location has been registered successfully. Please keep the location ID."
echo "Location ID = ${location_id}"
echo "${saved_filename} is created."
