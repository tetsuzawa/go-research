#!/bin/bash

observer_file="observer.json"

read -s -p 'Input your password: ' password1
echo ""

hash1=$(echo -n ${password1} | sha256sum | awk '{printf "%s", $1}')

unset password1

read -s -p 'Input your password, again: ' password2
echo ""

hash2=$(echo -n ${password2} | sha256sum | awk '{printf $1}')

unset password2

if [ ${hash1} != ${hash2} ]; then
    echo "Your two passwords do not match."
    exit 1
fi

unset hash2

cp ${observer_file} ${observer_file}.old
cat ${observer_file}.old | jq --arg hash1 ${hash1} '.password=$hash1' > ${observer_file}
rm -f ${observer_file}.old

echo "Your password is set in ${observer_file}."
