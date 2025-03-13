#!/bin/bash

if command -v git &>/dev/null; then
    git init
else
    echo "Git not found, skipping git initialization"
fi

go mod init $PROJ_NAME
