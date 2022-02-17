#!/bin/bash

if [[ -z $1 ]]; then
    echo "Specify an ent schema name"
    exit 1
fi

go run entgo.io/ent/cmd/ent init $1
