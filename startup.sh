#!/bin/bash

# requires go, postgresql, bash, and curl to be installed locally

# for windows
# start up your local db, the following is the default path for windows
pg_ctl -D "C:\Program Files\PostgreSQL\16\data" start

# start up postwoman, could also use 'go run server.go'
air

# for mac
