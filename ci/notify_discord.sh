#!/bin/sh
# This script is used to notify discord of post-build status

# ./ci/notify_discord.sh "webhook url" "Project Name" status
# success abort error failure


if [ "$3" == "success" ]
then
    color="2084461"
    message="build on $2 was successful"
elif [ "$3" == "abort" ]
then
    color="9131818"
    message="build on $2 was aborted"
elif [ "$3" == "error" ]
then
    color="16098851"
    message="build on $2 errored"
elif [ "$3" == "failure" ]
then
    color="15551285"
    message="build on $2 failed"
fi

timestamp=$(date --utc --iso-8601=seconds)
json="{\"embeds\":[{\"title\":\"Build on $2\",\"description\":\"$message\",\"url\":\"http://office.iosoftworks.net:8080/teams/main/pipelines/$2\",\"timestamp\":\"$timestamp\Z\",\"color\":\"$color\"}]}"

echo $json

curl -X POST --data "$json" --header "Content-Type:application/json" "$1"

