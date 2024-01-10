#!/usr/bin/env bash

# Install apt packages
sudo apt install silversearcher-ag entr

# Install go packages
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/a-h/templ/cmd/templ@latest
