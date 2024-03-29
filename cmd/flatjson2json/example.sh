#!/bin/sh

keys(){
    echo Timestamp
    echo SeverityText
}

j2j(){
    local cmd
    cmd=cat

    which jq  | fgrep -q jq && cmd='jq -c'
    which jaq | fgrep -q jaq && cmd='jaq -c'

    readonly cmd

    ${cmd}
}

export ENV_KEYS=$(
    keys |
        tr '\n' , |
        sed 's/,$//'
)

echo keys: ${ENV_KEYS}

echo '
{
    "Timestamp": "2024-01-18T14:25:22.494772866+09:00",
    "SeverityText": "INFO",
    "Body": "aufs is not supported",
    "Resource": {
        "service.name": "shoppingcart",
        "service.version": "3.1.4",
        "telemetry.sdk.language": "go"
    },
    "Attributes": {
        "server.address": "127.0.0.1"
    }
}
{
    "Timestamp": "2024-02-18T14:25:22.494772866+09:00",
    "SeverityText": "INFO",
    "Body": "aufs is not supported",
    "Resource": {
        "service.name": "shoppingcart",
        "service.version": "3.1.4",
        "telemetry.sdk.language": "go"
    },
    "Attributes": {
        "server.address": "127.0.0.1"
    }
}
' |
j2j |
./flatjson2json
