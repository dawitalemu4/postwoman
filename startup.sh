#!/bin/bash
# go into your postwoman dir
cd D:/developer/postwoman

# start up ypur local db
pg_ctl -D "C:\Program Files\PostgreSQL\16\data" start

# start up postwoman
air

# open browser 13234
# detect when chrome tab closed? and quit air 

# stop local db and kill air
pg_ctl -D "C:\Program Files\PostgreSQL\16\data" stop
kill -INT 888
