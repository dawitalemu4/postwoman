#!/bin/bash

# requires go, air, and postgresql to be installed locally

# create a shortcut with the target as the path to your terminal of choice and the start in as the path to parent directory of postwoman
# my target: "C:\Program Files\Git\bin\bash.exe" -i -l -c './postwoman/startup.sh'
# my start in path: "D:/developer"

# go into your postwoman dir
cd ./postwoman

# start up your local db
pg_ctl -D "C:\Program Files\PostgreSQL\16\data" start

# start up postwoman, could also use 'go run server.go'
air
