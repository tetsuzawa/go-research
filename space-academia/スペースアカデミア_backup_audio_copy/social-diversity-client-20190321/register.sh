#!/bin/bash

#server_address="http://localhost:8080"
server_address="https://social-diversity-devel.appspot.com"

observer_file="observer.json"
location_file="location.json"

wait_time=1

first_registration_response=$(curl -X POST -H "Content-Type: application/json" -d @${observer_file}  ${server_address}/accounts)

first_registration_responsee_body=$(echo ${first_registration_response} | jq -r '.error_message')

if [ "${first_registration_responsee_body}" != "null" ]; then
    echo ${first_registration_responsee_body}
    exit 1
fi

account_activation_token=$(echo ${first_registration_response} | jq -r ".account_activation_token")

sleep ${wait_time}

email=$(cat ${observer_file} | jq -r '.["e-mail"]')
password=$(cat ${observer_file} | jq -r '.password')

observer_id=$(curl -X PUT -H "Content-type: application/json" -d "{\"e-mail\":\"${email}\", \"password\":\"${password}\", \"token\": \"${account_activation_token}\"}" ${server_address}/accounts | jq -r ".observer_id")

sleep ${wait_time}

auth_success_token=$(curl -X POST -H "Content-Type: application/json" -d "{\"e-mail\": \"${email}\", \"password\": \"${password}\"}" ${server_address}/login | jq -r ".auth_success_token" )

if [ ${auth_success_token} = "" ]; then
    echo "Registering your account failed."
    exit 1
fi

saved_filename="${observer_file/./_${observer_id}.}"
cp ${observer_file} ${saved_filename}

cp ${location_file} ${location_file}.old
cat ${location_file}.old | jq -r --argjson observer_id ${observer_id} '.observer_id=$observer_id' > ${location_file}
rm -f ${location_file}.old

echo "Your account has been registered successfully. Please keep your observer ID."
echo "Observer ID = ${observer_id}"
echo "${saved_filename} is created."
